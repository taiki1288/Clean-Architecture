package database

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string ...interface{}) (Row, error)
}

type Result interface{
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}