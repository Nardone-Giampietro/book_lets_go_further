package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare a string containing the application version number
const version = "1.0.0"

// Struct that hold all the configuration settings for the application
type config struct {
	port int
	env  string
}

// Struct to hold the dependencies for the HTTP handlers, helpers and middleware.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Read the port number from the command line and use 8080 as default
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	app := &application{
		config: cfg,
		logger: logger,
	}

	r := app.routes()

	// Declare the HTTP server
	svr := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("starting %s server on port %d", cfg.env, cfg.port)
	if err := svr.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}

}
