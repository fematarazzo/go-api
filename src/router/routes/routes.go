package routes

import (
	"net/http"

	"api/src/middleware"
)

type Route struct {
	URI                   string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Configure(r *http.ServeMux) *http.ServeMux {
	routes := routesUsers
	routes = append(routes, routeLogin)

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.URI, middleware.Logger(middleware.Authenticate(route.Function)))
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Function))
		}
	}

	return r
}
