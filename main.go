package main

import (
	"flag"
	"log"
)

func main() {

	t := mustToken()

	//tgClient = telegram.New(token)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// cunsumer.Start(fetcher, processor) // Получение и обработка событий
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
