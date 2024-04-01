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
		RequireAuthentication: true,
	},
	{
		URI:                   "GET /users/{id}",
		Function:              controllers.ReadUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "PUT /users/{id}",
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "DELETE /users/{id}",
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
}
