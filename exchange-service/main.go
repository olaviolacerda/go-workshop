package main

import (
	"log"

	"github.com/olaviolacerda/exchange/internal/api"
)

func main() {
	log.Println("[INFO] Exchange Service")

	app := api.New()
	go app.Start()
	app.Stop()
}
