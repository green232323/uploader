package main

import (
	"fmt"
	"github.com/dnahurnyi/uploader/clientAPI/app/delivery"
	"github.com/dnahurnyi/uploader/clientAPI/app/handlers"
	"github.com/dnahurnyi/uploader/clientAPI/app/parse"
	"log"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

const (
	portDomainURL = "PORT_DOMAIN_URL"
	port          = "PORT"
)

func main() {
	ownPort := os.Getenv(port)
	if len(ownPort) == 0 {
		log.Fatal("can't get own port from env")
	}

	pdURL := os.Getenv(portDomainURL)
	if len(pdURL) == 0 {
		log.Fatal("Can't get port domain credentials from env")
	}

	deliveryClient, err := delivery.PortDomainClient(pdURL)
	if err != nil {
		log.Fatal("Cant connect to PortDomainService")
	}

	portHandler := handlers.PortsHandler{
		Log:          zerolog.New(os.Stdout),
		Parser:       parse.LargeJsonParser(),
		DomainClient: deliveryClient,
	}
	http.HandleFunc("/ports/", portHandler.Handle)
	fmt.Println("service started")
	http.ListenAndServe(":"+ownPort, nil)
}
