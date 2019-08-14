package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wexel-nath/wexel-email/pkg/api"
	"github.com/wexel-nath/wexel-email/pkg/config"
)

func main() {
	config.Configure()

	startServer()
}

func getListenAddress() string {
	port := config.GetPort()

	if len(port) == 0 {
		log.Fatal("PORT must be set")
	}

	return ":" + port
}

func startServer() {
	address := getListenAddress()
	fmt.Println("Listening on " + address)
	router := api.GetRouter()
	log.Fatal(http.ListenAndServe(address, router.GetHttpRouter()))
}
