package database

// database/user_repository.goでこれを呼び出す。
type SqlHandler interface {
	// 2つのメソッドの戻り値はResultとRowのインターフェースになっているのでこれによって依存ルールを守ることができる。
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

// Executeの戻り値
type Result interface{
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Queryの戻り値
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}