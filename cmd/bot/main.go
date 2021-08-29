package main

import (
	"log"
	"razanaubot/pkg/config"
	"razanaubot/pkg/data"
	"razanaubot/pkg/recurrent"
	"razanaubot/pkg/telegram"
)

func main() {
	dt, err := data.Get()
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

	rec := recurrent.BuildService(dt, tgService)

	if err := rec.StartWait(); err != nil {
		log.Fatal(err)
	}
}
