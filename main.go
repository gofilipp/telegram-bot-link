package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	token := mustToket()

	fmt.Println(token)
}

func mustToket() string {
	token := flag.String("t", "", "token for acces to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("broken token")
	}

	return *token
}
