package main

import (
	"flag"
	"fmt"
	"log"
	"tg-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustHost(), mustToket())

	fmt.Println(tgClient)
}

func mustHost() string {
	host := flag.String("h", tgBotHost, "host for acces to tg bot")

	flag.Parse()

	return *host
}

func mustToket() string {
	token := flag.String("t", "", "token for acces to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("broken token")
	}

	return *token
}
