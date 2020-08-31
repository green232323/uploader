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

func main() {
	deliveryClient, err := delivery.PortDomainClient(":9000")
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
	http.ListenAndServe(":8090", nil)
}
