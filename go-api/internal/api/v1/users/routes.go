package users

import (
	"api/internal/repos"
	"api/types"
)

func HandleRoutes(handlers map[string]func(gr repos.GlobalRepo, sessionUser *types.User, data []byte) (
	interface{}, error),
) map[string]func(gr repos.GlobalRepo, sessionUser *types.User, data []byte) (interface{}, error) {
	handlers["POST:Users"] = create

	return handlers
}
