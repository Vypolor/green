package app

import (
	"context"
	"log"
	"net/http"
)

type App struct {
	diContainer *diContainer
	serverMux   *http.ServeMux
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDIContainer,
		a.initServerMux,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDIContainer(_ context.Context) error {
	a.diContainer = newDIContainer()

	return nil
}

func (a *App) initServerMux(_ context.Context) error {
	mux := http.NewServeMux()
	a.diContainer.Handler().RegisterRoutes(mux)
	a.serverMux = mux

	return nil
}

func (a *App) runServer() error {
	serverAddr := a.diContainer.ServerConfig().Addr()
	log.Printf("server listening on: %s", serverAddr)

	return http.ListenAndServe(serverAddr, a.serverMux)
}
