package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	cfg := config{
		addr: ":8080",
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading ENV", "error", err)
	}

	// TODO: Database connection

	app := &application{
		config: cfg,
	}

	if err := app.run(app.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
