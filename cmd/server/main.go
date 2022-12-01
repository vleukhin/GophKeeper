package main

import (
	"context"
	"log"

	config "github.com/vleukhin/GophKeeper/internal/config/server"
	server "github.com/vleukhin/GophKeeper/internal/server/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	ctx := context.Background()

	server.Run(ctx, cfg)
}
