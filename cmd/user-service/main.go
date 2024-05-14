package main

import (
	"context"
	"github.com/EugeneTsydenov/go-user-service/internal/app"
)

func main() {
	ctx := context.Background()

	newApp, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	err = newApp.Serve()
	if err != nil {
		panic(err)
	}
}
