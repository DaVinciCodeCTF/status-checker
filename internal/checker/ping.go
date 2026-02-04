package checker

import (
	"time"

	"github.com/go-ping/ping"
)

type PingChecker struct {
	Name    string
	Target  string
	Timeout time.Duration
}

func (p *PingChecker) Check() CheckResult {
	pinger, err := ping.NewPinger(p.Target)
	if err != nil {
		return CheckResult{
			ServiceName: p.Name,
			Status:      "down",
			Message:     err.Error(),
			Timestamp:   time.Now(),
		}
	}

	pinger.Count = 3
	pinger.Timeout = p.Timeout
	pinger.SetPrivileged(true)

	start := time.Now()
	err = pinger.Run()
	latency := time.Since(start)

	if err != nil || pinger.Statistics().PacketsRecv == 0 {
		return CheckResult{
			ServiceName: p.Name,
			Status:      "down",
			Latency:     latency,
			Message:     "No packets received",
			Timestamp:   time.Now(),
		}
	}

	return CheckResult{
		ServiceName: p.Name,
		Status:      "up",
		Latency:     pinger.Statistics().AvgRtt,
		Message:     "ICMP ping successfull",
		Timestamp:   time.Now(),
	}
}
