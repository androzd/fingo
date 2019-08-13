package main

import (
	"github.com/androzd/finance/model"
	"github.com/androzd/finance/route"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"time"
)

func main() {
	model.Initialize()
	r := route.Initialize()

	go http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

	for {
		time.Sleep(1 * time.Minute)
	}
}
