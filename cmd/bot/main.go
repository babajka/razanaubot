package main

import (
	"log"
	"razanaubot/pkg/data"
	"razanaubot/pkg/recurrent"
	"razanaubot/pkg/telegram"
)

func main() {
	dt, err := data.Get()
	if err != nil {
		log.Fatal(err)
	}

	tgService := telegram.BuildService()

	rec := recurrent.BuildService(dt, tgService)

	if err := rec.StartWait(); err != nil {
		log.Fatal(err)
	}
}
