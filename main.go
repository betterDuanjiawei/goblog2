package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog2/app/http/middleware"
	"goblog2/bootstrap"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetRoute()

	http.ListenAndServe(":3000", middleware.RemoveTrailingSlash(router))

}
