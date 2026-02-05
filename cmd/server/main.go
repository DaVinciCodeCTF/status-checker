package main

import (
	"log"
	"os"

	"github.com/DaVinciCodeCTF/status-checker/internal/api"
	"github.com/DaVinciCodeCTF/status-checker/internal/config"
	"github.com/DaVinciCodeCTF/status-checker/internal/scheduler"
	"github.com/DaVinciCodeCTF/status-checker/internal/storage"
)

func main() {
	// Load config
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/services.yaml"
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Loaded config: %d services, check interval: %s, log level: %s", len(cfg.Services), cfg.CheckInterval, cfg.LogLevel)

	// Init storage
	store := storage.NewMemoryStorage()

	// Start scheduler in background
	sched := scheduler.New(cfg, store)
	go sched.Start()

	// Start API server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := api.NewServer(store, port)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
