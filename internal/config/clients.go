package config

import (
	"os"
	"strconv"
	"time"
)

const (
	defaultGreenApiBaseURL = "https://api.green-api.com"
	defaultCallTimeout     = 5 * time.Second
)

const (
	envGreenApiBaseURL         = "GREEN_API_BASE_URL"
	envGreenApiExternalTimeout = "GREEN_API_EXTERNAL_TIMEOUT"
)

type GreenAPIConfig struct {
	BaseURL string
	Timeout time.Duration
}

func NewGreenAPIConfig() (*GreenAPIConfig, error) {
	baseURL := os.Getenv(envGreenApiBaseURL)
	if baseURL == "" {
		baseURL = defaultGreenApiBaseURL
	}
	timeout := defaultCallTimeout
	if v := os.Getenv(envGreenApiExternalTimeout); v != "" {
		if sec, err := strconv.Atoi(v); err == nil {
			timeout = time.Duration(sec) * time.Second
		}
	}

	return &GreenAPIConfig{
		BaseURL: baseURL,
		Timeout: timeout,
	}, nil
}
