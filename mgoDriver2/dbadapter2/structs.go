package dbadapter2

import "github.com/globalsign/mgo"

type collections struct {
	collection map[string]*mgo.Collection
}

type session struct {
	session *mgo.Session
}

type mgoDriver struct {
	collection map[string]*mgo.Collection
	session    *mgo.Session
}

func (mgo mgoDriver) Do() {

}

func (mgo mgoDriver) Do1() {

}
