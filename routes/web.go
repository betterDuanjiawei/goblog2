package routes

import (
	"github.com/gorilla/mux"
	"goblog2/app/http/controllers"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {
	pc := new(controllers.PageController)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
}
