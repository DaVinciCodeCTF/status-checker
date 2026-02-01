package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Name           string            `yaml:"name"`
	Type           string            `yaml:"type"`
	Target         string            `yaml:"target"`
	Interval       time.Duration     `yaml:"interval,omitempty"`
	Timeout        time.Duration     `yaml:"timeout,omitempty"`
	ExpectedStatus int               `yaml:"expected_status,omitempty"`
	Headers        map[string]string `yaml:"headers,omitempty"`
}

type Config struct {
	CheckInterval time.Duration `yaml:"check_interval"`
	Timeout       time.Duration `yaml:"timeout"`
	LogLevel      string        `yaml:"log_level"`
	Services      []Service     `yaml:"services"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
