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
type Logged struct{
	Name string
	Empid int
	Role string
	Imgsrc string
}
type Prisoner struct{
	Id int
	Name string
	Ent string
	Exit string
	Cell string
	Charge string
}
type Employee struct{
	Name string
	Id int
	Dob string
	Role string
}
var Logged_user Logged
var user User
var userdb = "root"
var password = "anirudhdb"
var dab *sql.DB
var err error

func Queryer(query string, inp User) string {
	dab, err = sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer dab.Close()
	err = dab.QueryRow(query).Scan(&user.Password)
	switch err {
	case sql.ErrNoRows:
		return "User does not exist"
	default:
		if user.Password == inp.Password {
			return "OK"
		} else {
			return "Incorrect Password"
		}
	}
}

func Datahomepage(empid string)Logged{
	var y string
	var x []uint8

	dab, err = sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer dab.Close()
	query:="select * from employee where empid ="+empid+";"
	err = dab.QueryRow(query).Scan(&Logged_user.Name, &Logged_user.Empid, &x, &y, &Logged_user.Role)
	if err!=nil{
		panic(err)
	}
	dab.Close()
	return Logged_user
}

func Addprisonerquerier(pri_name string, pri_ent string, pri_ext string, pri_cell string, pri_charge string, query string){
	dab, err = sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer dab.Close()
	insert, err := dab.Prepare(query)
		if err != nil {
			panic(err)
		}
	 _,err = insert.Exec(pri_name, pri_ent, pri_ext, pri_cell, pri_charge)
	 if err != nil {
		panic(err)
	}
	insert.Close()
	dab.Close()
}

func Prisonertabledata()[]Prisoner{
	var prisoner Prisoner
	var prisoners []Prisoner
	dab, err = sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer dab.Close()
	query:="select * from prisoner;"
	data,err:=dab.Query(query)
	if(err!=nil){
		panic(err)
	}
	
	for data.Next(){
		data.Scan(&prisoner.Id,&prisoner.Name,&prisoner.Ent,&prisoner.Exit,&prisoner.Cell,&prisoner.Charge)
		prisoners=append(prisoners, prisoner)
	}
	data.Close()
	dab.Close()
	return prisoners
}

func Addemployeequerier(name string, dob string, address string, password1 string){
	dab, err = sql.Open("mysql", userdb+":"+password+"@/prisiondb")
	if err != nil {
		panic(err.Error())
	}
	defer dab.Close()
	query:="INSERT INTO `prisiondb`.`employee`(`name`,`dob`,`address`,`role`)VALUES(?,?,?,?);"
	println(name+dob+address+password1)
	insert, err := dab.Prepare(query)
		if err != nil {
			panic(err)
		}
	 _,err = insert.Exec(name,dob,address, "A")
	 if err != nil {
		panic(err)
	}
	insert.Close()
	query="INSERT INTO `prisiondb`.`password`(`password`)VALUES(?);"
	insert, err = dab.Prepare(query)
		if err != nil {
			panic(err)
		}
	 _,err = insert.Exec(password1)
	 if err != nil {
		panic(err)
	}
}

func Employeetabledata()[]Employee{
	var employee Employee
    var employees []Employee
    dab, err = sql.Open("mysql", userdb+":"+password+"@/prisiondb")
    if err != nil {
        panic(err.Error())
    }
    defer dab.Close()
    query:="select * from employee ;"
    data,err:=dab.Query(query)
    if(err!=nil){
        panic(err.Error())
    }
	defer data.Close()
    for data.Next(){
        data.Scan(&employee.Name,&employee.Id,&employee.Dob,&employee.Role)
        fmt.Println(employee.Id)
        employees=append(employees, employee)
    }
    return employees
}