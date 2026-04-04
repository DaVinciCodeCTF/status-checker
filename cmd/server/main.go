package main

import (
	"log"
	"os"

	"github.com/DaVinciCodeCTF/status-checker/internal/api"
	"github.com/DaVinciCodeCTF/status-checker/internal/config"
	"github.com/DaVinciCodeCTF/status-checker/internal/crypto"
	"github.com/DaVinciCodeCTF/status-checker/internal/scheduler"
	"github.com/DaVinciCodeCTF/status-checker/internal/storage"
	"github.com/joho/godotenv"
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

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	log.Printf("Loaded config: %d services, check interval: %s, log level: %s", len(cfg.Services), cfg.CheckInterval, cfg.LogLevel)

	// Encryption key (required)
	keyB64 := os.Getenv("ENCRYPTION_KEY")
	if keyB64 == "" {
		log.Fatal("ENCRYPTION_KEY is required (base64 of 32-byte AES key)")
	}

	encryptor, err := crypto.NewEncryptorFromBase64Key(keyB64)
	if err != nil {
		log.Fatalf("Invalid ENCRYPTION_KEY: %v", err)
	}

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

	server := api.NewServer(store, encryptor, port)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
