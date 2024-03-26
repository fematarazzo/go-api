package router

import (
	"net/http"

	"api/src/router/routes"
)

func Generate() *http.ServeMux {
	router := http.NewServeMux()

	return routes.Configure(router)
}
