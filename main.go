package main

import (
	"./api"
	"log"
	"os"
	"os/signal"
)

func main() {

	var (
		port      = os.Getenv("PORT")
		endpoints = os.Getenv("ENDPOINTS")
		host, _   = os.Hostname()
	)

	port = "3001"
	endpoints = "http://localhost:3001"

	app := api.AppSettings{
		Host:      host,
		Port:      port,
		Endpoints: endpoints,
	}

	log.Println("Starting tiny searcher service.")

	go api.StartServer(app)

	// block
	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, os.Interrupt, os.Kill)
	osSignal := <-osChan

	log.Printf("Tiny searcher service is exiting! OS signal: %v", osSignal)
}
