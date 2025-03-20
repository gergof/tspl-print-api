package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	configFilePath := flag.String("config", "", "Path of config file")
	listenAddr := flag.String("addr", "0.0.0.0:3000", "Address to listen on")
	flag.Parse()

	if *configFilePath == "" {
		fmt.Println("Error: Config flag is required!")
		flag.Usage()
		os.Exit(1)
	}

	log.Print("Starting TSPL print API")

	log.Printf("Loading config file from %s", *configFilePath)

	config, err := LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app := NewApp(config)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/ping", app.Ping)
	r.Post("/label/{endpoint}/render", app.Render)
	r.Post("/label/{endpoint}/print", app.Print)

	log.Printf("Start listening on %s", *listenAddr)
	if err := http.ListenAndServe(*listenAddr, r); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
