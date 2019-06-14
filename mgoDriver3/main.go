package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	// mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	db "github.com/go/mgoDriver3/dbadapter3"
)

type notification struct {
	ID               string   `json:"id,omitempty"`
	Type             string   `json:"type,omitempty"`
	Category         string   `json:"category,omitempty"`
	NotifyID         string   `json:"notifyID,omitempty"`
	JobID            string   `json:"jobID,omitempty"`
	CreatedBy        string   `json:"createdBy,omitempty"`
	CreatedByService string   `json:"createdByService,omitempty"`
	CreationDate     string   `json:"creationDate,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	UpdatedBy        string   `json:"updatedBy,omitempty"`
	ModifiedDate     string   `json:"modifiedDate,omitempty"`
	Notes            string   `json:"notes,omitempty"`
	Acknowledged     *bool    `json:"acknowledged,omitempty"`
	UserID           string   `json:"userid,omitempty" bson:"-"`
	Severity         string   `json:"severity,omitempty" bson:"-"`
	Summary          string   `json:"summary,omitempty" bson:"-"`
	Description      string   `json:"description,omitempty" bson:"-"`
	Remedy           string   `json:"remedy,omitempty" bson:"-"`
	PostTo           string   `json:"postTo,omitempty" bson:"-"`
}
type DbService struct {
	// Session    *mgo.Session
	// Collection *mgo.Collection
	db db.DBMongo
}

// InitSession initializes session
func (notifyMgo *DbService) InitSession() (err error) {
	// log.Debug("Entering")
	// defer log.Debug("Exiting")

	indexes := []string{"id", "category"}
	// log.Debug("Initilizig Db with indexes", indexes)
	notifyMgo.Session, notifyMgo.Collection, err = db.Init("notifications", indexes)
	return err

}

func (notifyMgo DbService) InsertNotification(notification notification) (string, string, error) {
	// log.Debug("Entering")
	// defer log.Debug("Exiting")

	if notifyMgo.Collection == nil {
		return "", "", errors.New("dbNotInitialized")
	}
	if notification.Type == "ALERT" {
		acknowledged := false
		notification.Acknowledged = &acknowledged
	}
	nowTime := int(time.Now().Unix())
	notification.CreationDate = strconv.Itoa(nowTime)
	notification.ID = bson.NewObjectId().Hex()
	err := notifyMgo.Collection.Insert(notification)
	if err != nil {
		// log.Error("error inserting ", notification, "error was: ", err.Error())
		return "", "", err
	}
	return notification.ID, notification.CreationDate, nil
}

func main() {
	instance := &DbService{}
	err := instance.InitSession()
	if err != nil {
		fmt.Println(err)
	}
	notification := notification{}
	_, _, err = instance.InsertNotification(notification)
	if err != nil {
		fmt.Println(err)
	}

}
