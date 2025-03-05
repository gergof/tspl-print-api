package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	configFilePath := flag.String("config", "", "Path of config file")
	flag.Parse()

	if *configFilePath == "" {
		log.Print("Config flag is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Print("Starting TSPL print API")

	log.Print("Loading config file from %s", configFilePath)

	config, err := LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	x, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(x))
}
