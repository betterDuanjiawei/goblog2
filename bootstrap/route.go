package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog2/routes"
)

func SetRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	return router
}
