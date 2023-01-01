package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	D"./lib"
)

var inp D.User

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
	}

	inp.Empid, _ = strconv.Atoi(r.FormValue("Empid"))
	inp.Password = r.FormValue("Password")

	query := "select password from password where empno = '" + r.FormValue("Empid") + "';"
	D.Queryer(query,inp)
}


var port = "8000"


func main() {
	

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	fmt.Println("Starting server at port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
