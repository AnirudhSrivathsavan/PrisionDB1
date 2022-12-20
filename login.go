package main

import (
	"database/sql"
	"fmt"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Empid    int    `json:"Empid"`
	Password string `json:"Password"`
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	//form parse shit
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
	}
	fmt.Fprintln(w, "sent data")

	/*
		empid := r.FormValue("empid")
		password := r.FormValue("password")


			Query to check if:
				1) given Empno exists
				2) if exists, does the password match
	*/
	empid := r.FormValue("empid")
	password := r.FormValue("password")
	var user User
	//query := "select password from password where empno=" + string(empid) + ";"
	query, err := db.Query("select password from password where empno = "+string(empid))
	if err != nil {
		fmt.Println("Some shit happened")
	}
	defer query.Close()
	fmt.Println(("******************\nempid = " + empid + "\tpassword = " + password))
	dis := db.QueryRow(password)
	err = dis.Scan(&user.Password)
	if err != nil {
		fmt.Println("Wrong Userid")
	}
	if user.Password != password {	
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("Granted")
}

var userdb = "root"
var password = "anirudhdb"
var port = "8000"
var db *sql.DB

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	//sql driver setup
	db, err := sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successful connection")

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/login", loginhandler)

	//server setup
	fmt.Println("Starting server at port " + port)
	log.Fatal(http.ListenAndServe(":8000", nil))

	//test zone
	fmt.Println(("Entered test zone"))

	fmt.Println(("Exited test zone"))
}
