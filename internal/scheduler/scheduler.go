package scheduler

import (
	"time"

	"github.com/DaVinciCodeCTF/status-checker/internal/config"
	"github.com/DaVinciCodeCTF/status-checker/internal/storage"
)

type Scheduler struct {
	config  *config.Config
	storage storage.Storage
}

func New(cfg *config.Config, store storage.Storage) *Scheduler {
	return &Scheduler{config: cfg, storage: store}
}

func (s *Scheduler) runChecks() {
	// TODO: Implement this function
	// It should loop through all services
	// and ping/curl (http request) them,
	// logging the results/status.
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(s.config.CheckInterval)
	defer ticker.Stop()

	// Immediate check
	s.runChecks()

	for range ticker.C {
		s.runChecks()
	}
}
