package main

//import (
//	"context"
//	"errors"
//	"github.com/go-kit/log"
//	"github.com/go-kit/log/level"
//	"io"
//	"os"
//	"os/signal"
//	"otus_project/internal/logger"
//	"otus_project/internal/service"
//	"syscall"
//)
//
//func main() {
//	loggerOutput := log.NewLogfmtLogger(os.Stdout)
//	loggerOutput = log.With(loggerOutput, "ts", log.DefaultTimestampUTC)
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	sigChan := make(chan os.Signal, 1)
//	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
//
//	go func() {
//		sig := <-sigChan
//		level.Info(loggerOutput).Log("msg", "Received shutdown signal", "signal", sig.String())
//		cancel()
//	}()
//
//	if err := logger.LoadAll(); err != nil {
//		if errors.Is(err, io.EOF) {
//			level.Info(loggerOutput).Log("msg", "No data to load", "err", err)
//		}
//		level.Error(loggerOutput).Log("msg", "Failed to load data", "err", err)
//	}
//
//	logger.StartSliceLogger(ctx, loggerOutput)
//
//	service.GenerateData(ctx)
//
//	level.Info(loggerOutput).Log("msg", "Program shut down successfully")
//}
