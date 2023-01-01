package lib

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

func Queryer(query string, inp User) {
	dab, err := sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer dab.Close()
	fmt.Println("Successful connection")
	err = dab.QueryRow(query).Scan(&user.Password)
	switch err {
	case sql.ErrNoRows:
		fmt.Printf("No entry")
	}
	print(user.Password)
}
