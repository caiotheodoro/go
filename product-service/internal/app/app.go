package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/caiotheodoro/product-service/internal/config"
	"github.com/caiotheodoro/product-service/internal/handler"
)

type App struct {
	config  config.Config
	handler http.Handler
}

func New(cfg config.Config) *App {
	app := &App{
		config:  cfg,
		handler: handler.NewHandler(cfg),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.handler,
	}

	fmt.Printf("Server running on port %d\n", a.config.ServerPort)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %s\n", err)
		}
	}()

	<-ctx.Done()

	return server.Shutdown(context.Background())
}