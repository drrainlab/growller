package main

import (
	"context"
	"fmt"
	"growwwler/internal/botservice"
	"growwwler/internal/config"
	"growwwler/internal/httpservice"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()

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

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done

	fmt.Println(cfg.BotConfig)

}
