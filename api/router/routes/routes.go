package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes representa todas as rotas da API
type Routes struct {
	Endpoint       string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// AllRoutes coloca todas as rotas dentro do router
func AllRoutes(r *mux.Router) *mux.Router {
	routesSlice := [][]Routes{
		userRoutes,
		serviceRoutes,
	}

	for _, routes := range routesSlice {
		for _, route := range routes {
			r.HandleFunc(
				route.Endpoint,
				route.Function,
			).Methods(route.Method)
		}
	}

	return r
}
