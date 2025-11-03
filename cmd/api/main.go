package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" // Import the pq driver for PostgreSQL
)

// Declare a string containing the application version number
const version = "1.0.0"

// Struct that hold all the configuration settings for the application
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
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

	// Set the database connection string based on the environment
	cfg.db.dsn = "postgres://greenlight:pa55word@localhost:5432/greenlight?sslmode=disable"

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// Set the database connection string based on the environment
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

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

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// create context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ping the database to check if it's reachable
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
