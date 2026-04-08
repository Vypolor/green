package config

import (
	"net"
	"os"
)

const (
	defaultServerHost = "localhost"
	defaultServerPort = "8080"
)

const (
	envServerHost = "SERVER_HOST"
	envServerPort = "SERVER_PORT"
)

type ServerConfig struct {
	Host string
	Port string
}

func (s *ServerConfig) Addr() string {
	return net.JoinHostPort(s.Host, s.Port)
}

func NewServerConfig() *ServerConfig {
	host := os.Getenv(envServerHost)
	if host == "" {
		host = defaultServerHost
	}
	port := os.Getenv(envServerPort)
	if port == "" {
		port = defaultServerPort
	}

	return &ServerConfig{
		Host: host,
		Port: port,
	}
}
