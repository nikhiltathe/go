package dbadapter

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
	// log.Debug("Entering")
	// defer log.Debug("Exiting")
	once.Do(func() {

		var mongoPort = ""
		// if os.Getenv("CRDEV") != "" {
		// 	mongoPort = u.EXTERNAL_STORPORT
		// } else {
		// 	mongoPort = u.INTERNAL_STORPORT
		// }

		// the service calls into this function need to open lockbox first
		// store := secureStore.GetSecureStore()
		// dbPw, err := store.RetrieveHashSecret(secureStore.MONGODBKEY)
		// if err != nil {
		// 	log.Error(err.Error())
		// }
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

// Init () - initializes the stor connection, returns the session and collections for users
func Init(collection string, indexes []string) (*mgo.Session, *mgo.Collection, error) {
	// log.Debug("Entering")
	// defer log.Debug("Exiting")
	dba = getDBAdapter()

	if dba.Session == nil {
		err := errors.New("Unable to connect to the CR db, Ensure CR db is running")
		return nil, nil, err
	}

	session := dba.Session.Copy()

	session.SetMode(mgo.Monotonic, true)

	col, err := AddCollection(session, collection, indexes)
	if err != nil {
		// log.Error(err)
		os.Exit(1)
	}
	return session, col, nil
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

// func VerifyPw(dbPw string) error {
// 	// log.Debug("Entering")
// 	// defer log.Debug("Exiting")

// 	log.Debug("connecting to:" + u.STORHOST + ":" + u.EXTERNAL_STORPORT)
// 	session, err := mgo.Dial(u.STORHOST + ":" + u.EXTERNAL_STORPORT)
// 	if err != nil {
// 		err = errors.New("Unable to connect to the db:" + err.Error())
// 		return err
// 	}
// 	defer session.Close()
// 	log.Debug("connected to:" + u.STORHOST + ":" + u.EXTERNAL_STORPORT)

// 	dbUser := session.DB(ADMIN)
// 	err = dbUser.Login(DBUSERNAME, dbPw)
// 	if err != nil {
// 		err = errors.New("Failed to login to admin db:" + err.Error())
// 		return err
// 	}
// 	return err
// }

// // This function changes both admin-db and cr-db pw with the same specified pw
// func ChangePw(newPw string) error {
// 	// log.Debug("Entering")
// 	// defer log.Debug("Exiting")

// 	log.Debug("connecting to:" + u.STORHOST + ":" + u.EXTERNAL_STORPORT)
// 	session, err := mgo.Dial(u.STORHOST + ":" + u.EXTERNAL_STORPORT)
// 	if err != nil {
// 		err = errors.New("Unable to connect to the db:" + err.Error())
// 		return err
// 	}
// 	defer session.Close()
// 	log.Debug("connected to:" + u.STORHOST + ":" + u.EXTERNAL_STORPORT)

// 	store := secureStore.GetSecureStore()
// 	err = store.OpenLockboxSys(true)
// 	defer store.CloseLockbox()
// 	if err != nil {
// 		err = errors.New("Unable to open lockbox:" + err.Error())
// 		return err
// 	}
// 	dbPw, err := store.RetrieveHashSecret(secureStore.MONGODBKEY)
// 	if err != nil {
// 		err = errors.New("Unable to retrieve DB credential lockbox:" + err.Error())
// 		return err
// 	}

// 	dbUser := session.DB(ADMIN)
// 	err = dbUser.Login(DBUSERNAME, dbPw)
// 	if err != nil {
// 		err = errors.New("Failed to login to admin db:" + err.Error())
// 		return err
// 	}
// 	log.Debug("logged into admin-db as:" + DBUSERNAME)

// 	updateDbUser := &mgo.User{
// 		Username: DBUSERNAME,
// 		Password: newPw,
// 	}
// 	err = dbUser.UpsertUser(updateDbUser)
// 	if err != nil {
// 		err = errors.New("Failed to change admin-db credential:" + err.Error())
// 		return err
// 	}
// 	dbUser.Logout()
// 	log.Debug("admin-db credential changed")

// 	crUser := session.DB("CR")
// 	err = crUser.Login(u.AuthUserName, dbPw)
// 	if err != nil {
// 		err = errors.New("Failed to login to cr db:" + err.Error())
// 		return err
// 	}
// 	log.Debug("logged into cr-db as:" + u.AuthUserName)

// 	updateCRUser := &mgo.User{
// 		Username: u.AuthUserName,
// 		Password: newPw,
// 	}
// 	err = crUser.UpsertUser(updateCRUser)
// 	if err != nil {
// 		err = errors.New("Failed to change cr-db credential:" + err.Error())
// 		return err
// 	}
// 	crUser.Logout()
// 	log.Debug("cr-db credential changed")

// 	return err
// }
