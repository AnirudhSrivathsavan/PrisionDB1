package lib2

import (
	"fmt"
	"os"
	D "login/lib1"
	"net/http"
	"strconv"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
)

type Status struct {
	Name	string
}

var inp D.User

func Formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
	}
	inp.Empid, _ = strconv.Atoi(r.FormValue("Empid"))
	inp.Password = r.FormValue("Password")
	query := "select password from password where empno = '" + r.FormValue("Empid") + "';"
	if D.Queryer(query, inp) {
		http.Redirect(w, r, "home.html", http.StatusSeeOther)
		fmt.Println("Im here")
	}else {
		td:=Status{"Wrong password!"}
		data, err := os.ReadFile("static/login.html")
		if err != nil {
			panic(err)
		}
		t, err := template.New("todos").Parse(string(data))
		if err != nil {
			panic(err)
		}
		err = t.Execute(w, td)
		if err != nil {
			panic(err)
		}
	}
}
