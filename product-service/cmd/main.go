package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/caiotheodoro/product-service/internal/app"
	"github.com/caiotheodoro/product-service/internal/config"
)

func main() {
	cfg := config.Load()
	application := app.New(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := application.Start(ctx)
	if err != nil {
		fmt.Println("Error starting application:", err)
		os.Exit(1)
	}
}