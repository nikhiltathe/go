package DB

import (
	log "github.com/go/gorillaMux/logger"
)

var users = map[string]string{
	"foo": "Mister Fooooo",
	"bar": "Missus Barrrr",
}

// Getuser returns user
func Getuser(id string) string {
	log.Debug("Entering")
	defer log.Debug("Exiting")

	return users[id]
}
