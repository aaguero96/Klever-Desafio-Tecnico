package routes

import (
	"net/http"

	user_controller "github.com/aaguero96/Klever-Desafio-Tecnico/api/controllers"
)

var userRoutes = []Routes{
	{
		Endpoint:       "/users",
		Method:         http.MethodPost,
		Function:       user_controller.Create,
		Authentication: false,
	},
	{
		Endpoint:       "/users",
		Method:         http.MethodGet,
		Function:       user_controller.Read,
		Authentication: false,
	},
	{
		Endpoint:       "/users/{userId}",
		Method:         http.MethodGet,
		Function:       user_controller.ReadById,
		Authentication: false,
	},
	{
		Endpoint:       "/users/{userId}",
		Method:         http.MethodPut,
		Function:       user_controller.Update,
		Authentication: false,
	},
}
