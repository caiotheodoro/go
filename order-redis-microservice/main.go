package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/caiotheodoro/go/order-redis-microservice/application"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println(err)
	}

}
