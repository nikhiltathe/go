package handler

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/go/gorillaMux/logger"
	"github.com/gorilla/mux"

	"github.com/go/gorillaMux/user/DB"
)

// UsersHandler retunrs producr name
func UsersHandler(w http.ResponseWriter, r *http.Request) {
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
	name := DB.Getuser(id)

	if name == "" {
		log.Error("No user found")
		json.NewEncoder(w).Encode("No ID found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	// Get corresponding order too
	order, err := getOrderDetails(id)
	if err != nil {
		log.Error(err)

		json.NewEncoder(w).Encode(name + " No order found")
	}
	json.NewEncoder(w).Encode(name + " order is :" + order)

}

func getOrderDetails(id string) (string, error) {
	log.Debug("Entering")
	defer log.Debug("Exiting")

	url := "http://127.0.0.1:9002/orders/" + id
	log.Debug("URL is : ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Error(err)
		return "", err
	}
	data, err := ioutil.ReadAll(res.Body)
	return string(data[:]), err
}
