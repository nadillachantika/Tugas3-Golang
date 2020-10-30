package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	cm "Tugas3/HtmlPage/common"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Index(w http.ResponseWriter, r *http.Request) {

	var employees []cm.Employees

	sql := `SELECT
				EmployeeID,
				IFNULL(LastName,''),
				IFNULL(FirstName,'') FirstName,
				IFNULL(Title,'') Title,
				IFNULL(TitleOfCourtesy,'') TitleOfCourtesy,
				IFNULL(BirthDate,'') BirthDate,
				IFNULL(HireDate,'') HireDate,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City ,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(HomePhone,'') HomePhone ,
				IFNULL(Extension,'') Extension,
				IFNULL(Photo,'') Photo,
				IFNULL(Notes,'') Notes
				
				
			FROM employees ORDER BY EmployeeID`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var employee cm.Employees
		err := result.Scan(&employee.EmployeeID, &employee.LastName, &employee.FirstName, &employee.Title,
			&employee.TitleOfCourtesy, &employee.BirthDate, &employee.HireDate, &employee.Address,
			&employee.City, &employee.Region, &employee.PostalCode, &employee.Country, &employee.HomePhone, &employee.Extension, &employee.Photo, &employee.Notes)

		if err != nil {
			panic(err.Error())
		}
		employees = append(employees, employee)
	}

	t, err := template.ParseFiles("index.html")
	t.Execute(w, employees)

	if err != nil {
		panic(err.Error())
	}

}

func main() {
	//<user>:<passwprd>@tcp<IP address>/<Password>
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)

}
