package DB

import (
	log "github.com/go/gorillaMux/logger"
)

var orders = map[string]string{
	"foo": "Pizza",
	"bar": "Burger",
}

// Getorder returns order
func Getorder(id string) string {
	log.Debug("Entering")
	defer log.Debug("Exiting")

	return orders[id]
}
