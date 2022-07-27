package routes

import (
	"net/http"

	upvote_controller "github.com/aaguero96/Klever-Desafio-Tecnico/api/controllers/upvote"
)

var upvoteRoutes = []Routes{
	{
		Endpoint:       "/upvotes",
		Method:         http.MethodPost,
		Function:       upvote_controller.Create,
		Authentication: false,
	},
	{
		Endpoint:       "/upvotes",
		Method:         http.MethodGet,
		Function:       upvote_controller.Read,
		Authentication: false,
	},
	{
		Endpoint:       "/upvotes/{upvoteId}",
		Method:         http.MethodGet,
		Function:       upvote_controller.ReadById,
		Authentication: false,
	},
	{
		Endpoint:       "/upvotes/{upvoteId}",
		Method:         http.MethodPut,
		Function:       upvote_controller.Update,
		Authentication: false,
	},
	{
		Endpoint:       "/upvotes/{upvoteId}",
		Method:         http.MethodDelete,
		Function:       upvote_controller.Delete,
		Authentication: false,
	},
}
