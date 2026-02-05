package scheduler

import (
	"log"
	"time"

	"github.com/DaVinciCodeCTF/status-checker/internal/checker"
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
	log.Println("Running checks...")

	for _, svc := range s.config.Services {
		go func(svc config.Service) {
			var chk checker.Checker

			timeout := svc.Timeout
			if timeout == 0 {
				timeout = s.config.Timeout
			}

			switch svc.Type {
			case "ping":
				chk = &checker.PingChecker{
					Name:    svc.Name,
					Target:  svc.Target,
					Timeout: timeout,
				}
			case "http":
				chk = &checker.HTTPChecker{
					Name:           svc.Name,
					Target:         svc.Target,
					Timeout:        timeout,
					ExpectedStatus: svc.ExpectedStatus,
					Headers:        svc.Headers,
				}
			default:
				log.Printf("Unknown checker type: %s", svc.Type)
				return
			}

			result := chk.Check()
			s.storage.Store(result)

			log.Printf("[%s] %s - %s (latency: %s)",
				result.Status, result.ServiceName, result.Message, result.Latency)
		}(svc)
	}
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
