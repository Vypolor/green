package main

import (
	"context"
	"green/internal/app"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize app: %s", err.Error())
	}
	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
