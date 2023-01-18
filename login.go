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
	http.HandleFunc("/form", R.Formhandler)
	http.HandleFunc("/form", R.Formhandler)
	fmt.Println("Starting server at port ", port)	
	fmt.Println("http://localhost:8000/login.html")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
