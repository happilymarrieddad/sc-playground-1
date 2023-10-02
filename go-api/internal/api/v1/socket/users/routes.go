package users

import (
	"api/internal/repos"
	"api/types"
)

func HandleRoutes(handlers map[string]func(gr repos.GlobalRepo, sessionUser *types.User, data []byte) *types.WSResponse,
) map[string]func(gr repos.GlobalRepo, sessionUser *types.User, data []byte) *types.WSResponse {
	handlers["POST:Users"] = create

	return handlers
}
