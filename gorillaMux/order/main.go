package main

import (
	"net/http"
	"os"

	log "github.com/go/gorillaMux/logger"
	"github.com/go/gorillaMux/order/handler"
	"github.com/gorilla/mux"
)

// Test : http://127.0.0.1:9002/orders/bar
// Test : http://127.0.0.1:9002/orders/foo

func main() {

	// logname string, level int, isStdOutput bool
	log.Init("logname", log.LEVEL_DEBUG, false)
	r := mux.NewRouter()

	r.HandleFunc("/orders/{id}", handler.OrdersHandler)
	r.HandleFunc("/", handler.OrdersHandler)
	log.Info("Starting server on port 9002")
	err := http.ListenAndServe("localhost:9002", r)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
