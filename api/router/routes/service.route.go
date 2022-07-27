package routes

import (
	"net/http"

	service_controller "github.com/aaguero96/Klever-Desafio-Tecnico/api/controllers/service"
)

var serviceRoutes = []Routes{
	{
		Endpoint:       "/services",
		Method:         http.MethodPost,
		Function:       service_controller.Create,
		Authentication: false,
	},
	{
		Endpoint:       "/services",
		Method:         http.MethodGet,
		Function:       service_controller.Read,
		Authentication: false,
	},
	{
		Endpoint:       "/services/{serviceId}",
		Method:         http.MethodGet,
		Function:       service_controller.ReadById,
		Authentication: false,
	},
	{
		Endpoint:       "/services/{serviceId}",
		Method:         http.MethodPut,
		Function:       service_controller.Update,
		Authentication: false,
	},
	{
		Endpoint:       "/services/{serviceId}",
		Method:         http.MethodDelete,
		Function:       service_controller.Delete,
		Authentication: false,
	},
}
