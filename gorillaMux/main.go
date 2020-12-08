package main

import (
	"net/http"
	"os"

	log "github.com/go/gorillaMux/logger"
	"github.com/gorilla/mux"
)

// Test : http://127.0.0.1:9001/users/bar
// Test : http://127.0.0.1:9001/users/foo

func main() {

	// logname string, level int, isStdOutput bool
	log.Init("logname", log.LEVEL_DEBUG, false)
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", handler.usersHandler)
	r.HandleFunc("/", handler.usersHandler)
	log.Info("Starting server on port 9001")
	err := http.ListenAndServe("localhost:9001", r)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
