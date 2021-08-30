package main

import (
	"context"
	"log"
	"razanaubot/pkg/config"
	"razanaubot/pkg/data"
	"razanaubot/pkg/db/redis"
	"razanaubot/pkg/telegram"
)

func main() {
	dataService, err := data.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	tgService, err := telegram.NewWithChannel(cfg.Telegram)
	if err != nil {
		log.Fatal(err)
	}

	redisService, err := redis.New(cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	lastTextID, err := redisService.GetLastSuccessfulTextID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	newText, newTextID := dataService.GetTextAfter(lastTextID)

	if err := tgService.SendMessage(newText); err != nil {
		log.Fatal(err)
	}

	log.Printf("text with ID '%v' was successfully sent to the channel", newTextID)

	if err := redisService.SetLastSuccessfulTextID(context.Background(), newTextID); err != nil {
		log.Fatal(err)
	}
}
