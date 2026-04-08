package app

import (
	greenapi "green/internal/clients/green-api"
	"green/internal/config"
	"green/internal/handler"
	"log"
	"sync"
)

type diContainer struct {
	serverConfigOnce   sync.Once
	serverConfig       *config.ServerConfig
	greenAPIConfigOnce sync.Once
	greenAPIConfig     *config.GreenAPIConfig

	greenAPIClientOnce sync.Once
	greenAPIClient     *greenapi.Client

	handlerOnce sync.Once
	handler     *handler.Handler
}

func newDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) ServerConfig() *config.ServerConfig {
	d.serverConfigOnce.Do(func() {
		d.serverConfig = config.NewServerConfig()
	})

	return d.serverConfig
}

func (d *diContainer) GreenAPIConfig() *config.GreenAPIConfig {
	d.greenAPIConfigOnce.Do(func() {
		cfg, err := config.NewGreenAPIConfig()
		if err != nil {
			log.Fatalf("failed to get green-api config: %v", err)
		}
		d.greenAPIConfig = cfg
	})

	return d.greenAPIConfig
}

func (d *diContainer) GreenAPIClient() *greenapi.Client {
	d.greenAPIClientOnce.Do(func() {
		d.greenAPIClient = greenapi.NewClient(d.GreenAPIConfig())
	})

	return d.greenAPIClient
}

func (d *diContainer) Handler() *handler.Handler {
	d.handlerOnce.Do(func() {
		d.handler = handler.New(d.GreenAPIClient())
	})

	return d.handler
}
