package mysql

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

func GetHandle(user, pw, logicalDB string) (db *sql.DB, err error) {
    db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(10.2.0.2:3306)/%s", user, pw, logicalDB))
    return
}
