package lib1

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Empid    int    `json:"Empid"`
	Password string `json:"Password"`
}

var user User
var userdb = "root"
var password = "anirudhdb"
var dab *sql.DB
var err error

func Queryer(query string, inp User) bool{
	dab, err = sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer dab.Close()
	err = dab.QueryRow(query).Scan(&user.Password)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("User does not exist")
	default:
		if user.Password == inp.Password {
			return true
		} else {
			return false
		}
	}
	return false
}
