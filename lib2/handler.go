package lib2

import (
	"fmt"
	"html/template"
	D "login/lib1"
	"net/http"
	"os"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type Status struct {
	Name string
}

var Logged_user D.Logged
var Session bool
var inp D.User

func Formhandler(w http.ResponseWriter, r *http.Request) {
	if err :=
		r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
	}
	inp.Empid, _ = strconv.Atoi(r.FormValue("Empid"))
	inp.Password = r.FormValue("Password")
	query := "select password from password where empid= '" + r.FormValue("Empid") + "';"
	x := D.Queryer(query, inp)
	if x == "OK" {
		Session = true
		http.Redirect(w, r, "home.html", http.StatusSeeOther)
	} else {
		td := Status{x}
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

func Loginpage(w http.ResponseWriter, r *http.Request) {
	td := Status{""}
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

func Homepage(w http.ResponseWriter, r *http.Request) {
	if !Session {
		http.Redirect(w, r, "login.html", http.StatusSeeOther)
	}
	Logged_user = D.Datahomepage(strconv.Itoa(inp.Empid))
	data, err := os.ReadFile("static/home.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("todos").Parse(string(data))
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, Logged_user)
	if err != nil {
		panic(err)
	}
}

func Addprisonerpage(w http.ResponseWriter, r *http.Request) {
	/*
	if !Session {
		http.Redirect(w, r, "login.html", http.StatusSeeOther)
	}
	*/
	data, err := os.ReadFile("static/addprisoner.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("todos").Parse(string(data))
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, Logged_user)
	if err != nil {
		panic(err)
	}
}

func Addprisonerform(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
	}
	pri_name := r.FormValue("name")
	pri_ent := r.FormValue("entdate")
	pri_ext := r.FormValue("extdate")
	pri_cell := r.FormValue("cell")
	pri_charge := r.FormValue("charge")
	query := "insert into prisoner (`pri_name`,`ent_date`,`ext_date`,`cell`,`charge`)values(?,?,?,?,?);"
	
	fmt.Println(pri_name, pri_ent, pri_ext, pri_cell, pri_charge)
	D.Addprisonerquerier(pri_name, pri_ent, pri_ext, pri_cell, pri_charge,query)
	Homepage(w,r)	
}

func Prisonertable(w http.ResponseWriter, r *http.Request){
	prisoners:=D.Prisonertabledata()
	data, err := os.ReadFile("static/prisonertable.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("todos").Parse(string(data))
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, prisoners)
	if err != nil {
		panic(err)
	}
}

func Addemployeepage(w http.ResponseWriter, r *http.Request) {
	/*
	if !Session {
		http.Redirect(w, r, "login.html", http.StatusSeeOther)
	}
	*/
	data, err := os.ReadFile("static/addemployee.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("todos").Parse(string(data))
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, Logged_user)
	if err != nil {
		panic(err)
	}
}

func Addemployee(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
	}
	pri_name := r.FormValue("name")
	pri_ent := r.FormValue("dob")
	pri_cell := r.FormValue("Address")
	pri_charge := r.FormValue("Password")	
	D.Addemployeequerier(pri_name, pri_ent, pri_cell, pri_charge)
	Homepage(w,r)	
}

func Employeetable(w http.ResponseWriter, r *http.Request){
	employees:=D.Employeetabledata()
	data, err := os.ReadFile("static/employeetable.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("todos").Parse(string(data))
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, employees)
	if err != nil {
		panic(err)
	}
}