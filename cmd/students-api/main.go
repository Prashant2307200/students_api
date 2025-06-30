package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Prashant2307200/students-api/internal/config"
	"github.com/Prashant2307200/students-api/internal/http/handlers/student"
)

func main() {

	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	server := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.HttpServer.Addr))
	fmt.Printf("Server is running... %s", cfg.HttpServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server.")
		}
	}()
	<-done // receive signal as block to complete present task

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully.")
}
