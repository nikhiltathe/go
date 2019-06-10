package dbadapter2

type dbCollection interface {
	Do()
}
type dbSession interface {
	Do1()
}

type DBMongo interface {
	dbCollection
	dbSession
}
