package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading ENV file!")
	}
	cfg := config{
		addr: ":8080",
	}

	// TODO: Database connection

	app := &application{
		config: cfg,
	}

	if err := app.run(app.mount()); err != nil {
		log.Panicf("Server failed to connect. Error %s", err)
		os.Exit(1)
	}
}
