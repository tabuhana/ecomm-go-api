package main

import "os"

func main() {
	cfg := config{
		addr: ":8080",
	}

	// TODO: Logger

	// TODO: Database connection

	app := &application{
		config: cfg,
	}

	if err := app.run(app.mount()); err != nil {
		os.Exit(1)
	}
}
