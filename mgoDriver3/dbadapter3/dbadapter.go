package dbadapter3

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/globalsign/mgo" // our DB driver
)

const ADMIN = "admin"
const DBUSERNAME = "dbadmin"

type DBAdapter struct {
	Session *mgo.Session // our cloneable session
}

var dba *DBAdapter
var once sync.Once

// GetStorAdapter - gets the singleton for CR Persistent Stor
func getDBAdapter() *DBAdapter {
	once.Do(func() {

		var mongoPort = "27017"
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{"mongo-auth" + ":" + mongoPort},
			Timeout:  10 * time.Second,
			Username: "crdbadmin",
			Password: "fkVOdjC87dhEJIbV",
			Database: "CR",
		}

		session, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			// log.Error(err.Error())
		}
		dba = &DBAdapter{Session: session}
	})
	return dba
}

// // Init () - initializes the stor connection, returns the session and collections for users
// func Init(collection string, indexes []string) (*mgo.Session, *mgo.Collection, error) {
// 	// log.Debug("Entering")
// 	// defer log.Debug("Exiting")
// 	dba = getDBAdapter()

// 	if dba.Session == nil {
// 		err := errors.New("Unable to connect to the CR db, Ensure CR db is running")
// 		return nil, nil, err
// 	}

// 	session := dba.Session.Copy()

// 	session.SetMode(mgo.Monotonic, true)

// 	col, err := AddCollection(session, collection, indexes)
// 	if err != nil {
// 		// log.Error(err)
// 		os.Exit(1)
// 	}
// 	return session, col, nil
// }

// Init () - initializes the stor connection, returns the session and collections for users
func Init(collection string, indexes []string) (DBMongo, error) {
	// log.Debug("Entering")
	// defer log.Debug("Exiting")
	dba = getDBAdapter()

	if dba.Session == nil {
		err := errors.New("Unable to connect to the CR db, Ensure CR db is running")
		return nil, err
	}

	session := dba.Session.Copy()

	session.SetMode(mgo.Monotonic, true)

	col, err := AddCollection(session, collection, indexes)
	if err != nil {
		// log.Error(err)
		os.Exit(1)
	}

	mgo := mgoDriver{}
	mgo.collection[collection] = col
	mgo.session = session
	// return session, col, nil
	return mgo, nil
}

func AddCollection(session *mgo.Session, collection string, indexes []string) (*mgo.Collection, error) {
	// log.Debug("Entering")
	// defer log.Debug("Exiting")

	// opening the collection for DB "ir"
	col := session.DB("CR").C(collection)

	index := mgo.Index{
		Key:    indexes,
		Unique: true,
	}

	err := col.EnsureIndex(index)
	return col, err
}

// altInit () - initializes the stor connection, returns the session and collections for users
func altInit(collection string, indexes []string) (*mgo.Session, *mgo.Collection, error) {

	dba = getDBAdapter()

	if dba.Session == nil {
		err := errors.New("Unable to connect to the CR db, Ensure CR db is running")
		return nil, nil, err
	}

	session := dba.Session.Copy()

	session.SetMode(mgo.Monotonic, true)

	// opening the collection for DB "ir"
	col := session.DB("CR").C(collection)

	idIndexStr := []string{"id"}
	idIndex := mgo.Index{Key: idIndexStr, Unique: true}

	err := col.EnsureIndex(idIndex)
	if err != nil {
		return session, col, err
	}

	if len(indexes) > 0 {
		// now unique index by the callers request. the set of indexes make it unique
		colIndexes := mgo.Index{
			Key:    indexes,
			Unique: true,
		}

		err = col.EnsureIndex(colIndexes)
	}

	return session, col, err
}
