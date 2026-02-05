package storage

import (
	"sync"

	"github.com/DaVinciCodeCTF/status-checker/internal/checker"
)

type MemoryStorage struct {
	mu      sync.RWMutex
	results map[string]checker.CheckResult
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		results: make(map[string]checker.CheckResult),
	}
}

func (m *MemoryStorage) Store(result checker.CheckResult) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.results[result.ServiceName] = result
}

func (m *MemoryStorage) GetAll() map[string]checker.CheckResult {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Copy to avoid race conditions
	copy := make(map[string]checker.CheckResult)
	for k, v := range m.results {
		copy[k] = v
	}
	return copy
}
