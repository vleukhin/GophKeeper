package main

import (
	"context"
	config "github.com/vleukhin/GophKeeper/internal/config/server"
	server "github.com/vleukhin/GophKeeper/internal/server/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	ctx := context.Background()

	server.Run(ctx, cfg)
}
