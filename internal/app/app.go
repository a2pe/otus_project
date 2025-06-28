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
)

type App struct {
	cfg *config.Config
	ctx context.Context
}

func NewApp(ctx context.Context) (*App, error) {
	return &App{
		cfg: config.NewConfig(),
		ctx: ctx,
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
	})
}

func (a *App) Start() error {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := chi.NewRouter()
	a.registerRoutes(router)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	log.Println("Server is running...")

	<-ctx.Done()

	log.Println("Shutting down server gracefully...")
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxTimeout); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server exited cleanly")
	return nil
}

func (a *App) Stop() error {
	// опционально, если захочешь вручную завершать App
	return nil
}
