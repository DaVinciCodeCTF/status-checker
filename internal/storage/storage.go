package storage

import "github.com/DaVinciCodeCTF/status-checker/internal/checker"

type Storage interface {
	Store(result checker.CheckResult)
	GetAll() map[string]checker.CheckResult
}
