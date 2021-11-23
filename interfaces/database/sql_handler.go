package database

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string ...interface{}) (Row, error)
}