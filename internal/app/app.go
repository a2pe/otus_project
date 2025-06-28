// @title Productivity Tracker API
// @version 1.0
// @description API for tracking users, projects, tasks and more.
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"

	"os/signal"
	"otus_project/internal/config"
	"otus_project/internal/handler"
	"syscall"
	"time"

	"github.com/swaggo/http-swagger"
	_ "otus_project/docs" // важно для init() из swag
)

type App struct {
	cfg *config.Config
	ctx context.Context
}

func NewApp(ctx context.Context) (*App, error) {
	return &App{
		config.NewConfig(),
		ctx,
	}, nil
}

func (a *App) registerRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		itemTypes := []string{"user", "project", "task", "reminder", "tag", "time_entry"}
		for _, t := range itemTypes {
			r.Get(fmt.Sprintf("/%s", t), handler.GetAllHandler(t))
			r.Post(fmt.Sprintf("/%s", t), handler.CreateItemHandler(t))
			r.Put(fmt.Sprintf("/%s/{id}", t), handler.UpdateItemHandler(t))
			r.Delete(fmt.Sprintf("/%s/{id}", t), handler.DeleteItemHandler(t))
			r.Get(fmt.Sprintf("/%s/{id}", t), handler.GetItemByIDHandler(t))
		}
		r.Get("/swagger/*", httpSwagger.WrapHandler)

	})
}

func (a *App) Start() error {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	router := chi.NewRouter()
	a.registerRoutes(router)

	serverHTTP := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler: router,
	}

	go func() {
		if err := serverHTTP.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	//<-a.ctx.Done() // А в чем разница между такими вызовами? Вызов завершенного контекста в App или это он же и есть?
	//stop()

	log.Println("Shutting down server")
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := serverHTTP.Shutdown(ctxTimeout); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server Exited Properly")
	return nil
}

func (a *App) Stop() error {
	return nil
}
