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
}
