package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/go/gorillaMux/logger"
	"github.com/gorilla/mux"

	"github.com/go/gorillaMux/order/DB"
)

// OrdersHandler retunrs producr name
func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Entering")
	defer log.Debug("Exiting")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		log.Error("ID missing")
		json.NewEncoder(w).Encode("ID missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Info("Searching for ", id)
	name := DB.Getorder(id)

	if name != "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(name)
		return
	}

	json.NewEncoder(w).Encode("No ID found")
	w.WriteHeader(http.StatusNotFound)

}
