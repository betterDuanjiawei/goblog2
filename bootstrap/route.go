package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog2/pkg/route"
	"goblog2/routes"
)

func SetRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)
	return router
}
