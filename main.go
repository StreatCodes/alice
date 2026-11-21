package main

import (
	"log"
	"net/http"
)

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config %s\n", err)
	}

	http.Handle("/", proxyHandler(config))

	log.Printf("Listening on %s", config.ListenAddress)
	log.Fatal(http.ListenAndServe(config.ListenAddress, nil))
}
