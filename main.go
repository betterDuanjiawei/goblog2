package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog2/bootstrap"
	"goblog2/pkg/database"
	"goblog2/pkg/logger"
	"net/http"
	"strconv"
	"strings"
)

var router *mux.Router
var db *sql.DB

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetRoute()

	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("POST").Name("articles.delete")
	// 自定义404页面

	// 中间件,强制 Content-Type 类型为 text/html
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))

}

func getArticleByID(id string) (Article, error) {
	query := "SELECT * FROM ARTICLES WHERE id = ?"
	article := Article{}
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)

	return article, err
}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置标头
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}

type Article struct {
	Title, Body string
	ID          int64
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func articlesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := getRouteVariable("id", r)
	article, err := getArticleByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "400 未找到文章")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		rowsAffected, err := article.Delete()
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			if rowsAffected > 0 {
				indexURL, _ := router.Get("articles.index").URL()
				http.Redirect(w, r, indexURL.String(), http.StatusFound) // 这里必须是302跳转
			} else {
				// 有可能并发,被删除了
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "400 文章未找到")
			}
		}

	}
}

func (a Article) Delete() (rowsAffected int64, err error) {
	// ID 是数据库查出来的,无需担心 sql注入,用纯文本格式,减少一次 sql请求
	rs, err := db.Exec("DELETE FROM articles WHERE id = " + strconv.FormatInt(a.ID, 10))
	if err != nil {
		return 0, err
	}
	if n, _ := rs.RowsAffected(); n > 0 {
		return n, nil
	}
	return 0, nil
}
