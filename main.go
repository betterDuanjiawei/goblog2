package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog2/bootstrap"
	"goblog2/pkg/database"
	"net/http"
	"strings"
)

var router *mux.Router
var db *sql.DB

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetRoute()

	// 自定义404页面

	// 中间件,强制 Content-Type 类型为 text/html
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))

}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置标头
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}
