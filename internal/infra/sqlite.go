package infra

import (
	"database/sql"
	"fmt"

	sqlite3 "github.com/ncruces/go-sqlite3"
)

type SqliteDB struct {
	*sql.DB
}

func ConnectToSqliteDB() {
	db, openErr := sql.Open("sqlite3", fmt.Sprintf("file:%s", appConfig.DBEnv))
	
}