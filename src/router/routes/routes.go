package routes

import "net/http"

type Route struct {
	URI                   string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Configure(r *http.ServeMux) *http.ServeMux {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function)
	}

	return r
}
