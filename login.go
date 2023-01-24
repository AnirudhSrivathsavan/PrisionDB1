package main

import (
	"fmt"
	"log"
	"net/http"
	R "login/lib2"
)

var port = "8000"

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/login.html", R.Loginpage)
	http.HandleFunc("/home.html", R.Homepage)
	http.HandleFunc("/addprisoner.html", R.Addprisonerpage)
	http.HandleFunc("/addemployee.html", R.Addemployeepage)
	http.HandleFunc("/prisonertable.html", R.Prisonertable)
	http.HandleFunc("/employeetable.html", R.Employeetable)
	http.HandleFunc("/addprisoner",R.Addprisonerform)
	http.HandleFunc("/addemployee", R.Addemployee)
	http.HandleFunc("/form", R.Formhandler)
	fmt.Println("Starting server at port ", port)	
	fmt.Println("http://localhost:8000/login.html")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
