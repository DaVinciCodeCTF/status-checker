package checker

import "time"

type CheckResult struct {
	ServiceName string        `json:"service_name"`
	Status      string        `json:"status"`
	Latency     time.Duration `json:"latency"`
	Message     string        `json:"message"`
	Timestamp   time.Time     `json:"timestamp"`
}

type Checker interface {
	Check() CheckResult
}
