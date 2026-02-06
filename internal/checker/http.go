package checker

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPChecker struct {
	Name           string
	Target         string
	Timeout        time.Duration
	ExpectedStatus int
	Headers        map[string]string
}

func (h *HTTPChecker) Check() CheckResult {
	client := &http.Client{Timeout: h.Timeout}

	req, err := http.NewRequest("GET", h.Target, nil)
	if err != nil {
		return CheckResult{
			ServiceName: h.Name,
			Status:      "down",
			Message:     err.Error(),
			Timestamp:   time.Now(),
		}
	}

	for k, v := range h.Headers {
		req.Header.Set(k, v)
	}

	start := time.Now()
	resp, err := client.Do(req)
	latency := time.Since(start)

	if err != nil {
		return CheckResult{
			ServiceName: h.Name,
			Status:      "down",
			Latency:     latency,
			Message:     err.Error(),
			Timestamp:   time.Now(),
		}
	}
	defer resp.Body.Close()

	status := "up"
	message := fmt.Sprintf("HTTP %d", resp.StatusCode)

	if h.ExpectedStatus != 0 && resp.StatusCode != h.ExpectedStatus {
		status = "degraded"
		message = fmt.Sprintf("Expected %d, got %d", h.ExpectedStatus, resp.StatusCode)
	}

	return CheckResult{
		ServiceName: h.Name,
		Status:      status,
		Latency:     latency,
		Message:     message,
		Timestamp:   time.Now(),
	}
}
