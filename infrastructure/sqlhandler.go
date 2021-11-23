package infrastructure

// DBは外部パッケージを使うので一番外側のinfrastructure層になる。
import (
	"Clean-Architecture/interfaces/database"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


type SqlHandler struct {
	// *Sql.DBはdatabase/sqlパッケージのDB接続に必要なtype
	Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/CleanArchitecture")
	if err != nil {
		panic(err.Error)
	}
	// newは指定した方のポインタ型を生成する関数。(構造体の初期化)
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}