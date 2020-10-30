package main

import (
	"database/sql"
	"fmt"
	"net/http"
	fn "pnp/HtmlPostData/function"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//init router
	r := mux.NewRouter()

	http.HandleFunc("/", fn.RouteIndexGet)
	http.HandleFunc("/process", fn.RouteSubmitPost)
	r.HandleFunc("/process", fn.RouteSubmitPost).Methods("POST")

	fmt.Println("server started at localhost:9090")
	http.ListenAndServe(":9090", r)

}
