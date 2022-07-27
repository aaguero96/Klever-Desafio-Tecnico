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
}
