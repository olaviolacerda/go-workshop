package main

import (
	"github.com/olaviolacerda/account/internal/infra"
	"log"

	"github.com/olaviolacerda/account/internal/api"
)

func main() {
	log.Println("[INFO] Account Service")

	infra.InitDB()

	app := api.New()
	go app.Start()
	app.Stop()
}
