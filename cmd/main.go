package main

import (
	"context"
	"fmt"
	"growwwler/internal/botservice"
	"growwwler/internal/config"
	"growwwler/internal/httpservice"
)

func main() {

	ctx := context.Background()

	done := make(chan bool)

	cfg, err := config.New(ctx)
	if err != nil {
		panic(err)
	}
	// bot service

	tgService, err := botservice.NewBotService(&cfg.BotConfig)
	if err != nil {
		panic(err)
	}

	srv, err := httpservice.NewHTTPService(&cfg.HttpService, tgService)
	if err != nil {
		panic(err)
	}

	go srv.Run(ctx, &cfg.HttpService)

	err = tgService.Run(ctx)
	if err != nil {
		panic(err)
	}

	<-done

	fmt.Println(cfg.BotConfig)

}
