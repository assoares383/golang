package main

import (
	"encurtador-url/api"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("All systems offline!")
}

func run() error {
	db := make(map[string]string)

	hander := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      hander,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
