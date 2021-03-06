package dao

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

var (
	ErrorNotRows = errors.New("There is no data.")
)

// user info
type User struct {
	ID   int
	Name string
}

func GetUserByID(id int) (string, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err != nil {
		return "", xerrors.Wrap(err, "db conneted faild.")
	}
	var name string
	err = db.QueryRow("Select id, name from user where id=?", id).Scan(&name)
	if err != nil {
		return "", xerrors.Wrap(err, "search failed")
	}
	return name, nil
}
