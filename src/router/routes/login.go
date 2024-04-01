package routes

import "api/src/controllers"

var routeLogin = Route{
	URI:                   "POST /login",
	Function:              controllers.Login,
	RequireAuthentication: false,
}
