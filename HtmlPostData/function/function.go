package function

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

var db *sql.DB
var err error

func RouteIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func RouteSubmitPost(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	if r.Method == "POST" {

		var tmpl = template.Must(template.New("result").ParseFiles("index.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var firstname = r.FormValue("firstname")
		var lastname = r.Form.Get("lastname")

		var data = map[string]string{"firstname": firstname, "lastname": lastname}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//tugas insertkan ke database ke table user

		return
	}
	http.Error(w, "", http.StatusBadRequest)

	if r.Method == "POST" {
		EmployeeID := r.FormValue("EmployeeID")
		LastName := r.FormValue("LastName")
		FirstName := r.FormValue("FirstName")
		Title := r.FormValue("Title")
		TitleOfCoutesy := r.FormValue("TitleOfCourtesy")
		BirthDate := r.FormValue("BirthDate")
		HireDate := r.FormValue("HireDate")
		Adress := r.FormValue("Address")
		City := r.FormValue("City")
		Region := r.FormValue("Region")
		PostalCode := r.FormValue("PostalCode")
		Country := r.FormValue("Country")
		HomePhone := r.FormValue("HomePhone")
		Extension := r.FormValue("Extension")
		Photo := r.FormValue("Photo")
		Notes := r.FormValue("Notes")

		stmt, err := db.Prepare("INSERT INTO employees (EmployeeID,LastName,FirstName,Title,TitleOfCourtesy,BirthDate,HireDate,Address,City,Region,PostalCode,HomePhone,Extension,Photo,Notes) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(EmployeeID, LastName, FirstName, Title, TitleOfCoutesy, BirthDate, HireDate, Adress, City, Region, PostalCode, Country, HomePhone, Extension, Photo, Notes)
		if err != nil {
			fmt.Println(w, "Data Duplicate")
		} else {
			fmt.Println(w, "Data Created")
		}

	}

}
