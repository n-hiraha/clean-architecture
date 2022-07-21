package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


//mysqlとのコネクション
type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:password@tcp(nikki-mysql:3306)/nikki")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}