package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/olaviolacerda/exchange/internal/infra"
)

type App struct {
	server *http.Server
}

func New() App {
	host := infra.GetConfig("host")
	port := infra.GetConfig("port")

	routes := configureRoutes()

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Handler:      routes,
		IdleTimeout:  15 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return App{srv}
}

func (a App) Start() {
	log.Printf("[INFO] Server is ready on http://%s\n", a.server.Addr)

	err := a.server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("[ERROR] Cannot initiate the server, reason: %s \n", err.Error())
	}
}

func (a App) Server() *http.Server {
	return a.server
}

func (a App) Stop() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	<-ctx.Done()

	log.Println("[INFO] Server is shutting down...")
	// stop receiving any request.
	if err := a.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("[FAIL] Fail on stop the server, reason: %s\n", err.Error())
	}
	log.Println("[INFO] Server has been stopped.")
	stop()
}
