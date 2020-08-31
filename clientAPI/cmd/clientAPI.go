package main

import (
	"fmt"
	"github.com/dnahurnyi/uploader/clientAPI/app/handlers"
	"github.com/dnahurnyi/uploader/clientAPI/app/parse"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	portHandler := handlers.PortsHandler{
		Log:    zerolog.New(os.Stdout),
		Parser: parse.LargeJsonParser(),
	}
	http.HandleFunc("/ports/", portHandler.Handle)
	fmt.Println("service started")
	http.ListenAndServe(":8090", nil)
}
