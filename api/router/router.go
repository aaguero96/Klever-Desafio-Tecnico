package router

import (
	"github.com/aaguero96/Klever-Desafio-Tecnico/api/router/routes"
	"github.com/gorilla/mux"
)

// Router gera um router com as rotas configuradas
func Router() *mux.Router {
	router := mux.NewRouter()
	return routes.AllRoutes(router)
}
