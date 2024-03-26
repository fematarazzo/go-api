package routes

import (
	"api/src/controllers"
)

var routesUsers = []Route{
	{
		URI:                   "POST /users",
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "GET /users",
		Function:              controllers.ReadUsers,
		RequireAuthentication: false,
	},
	{
		URI:                   "GET /users/{id}",
		Function:              controllers.ReadUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "PUT /users/{id}",
		Function:              controllers.UpdateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "DELETE /users/{id}",
		Function:              controllers.DeleteUser,
		RequireAuthentication: false,
	},
}
