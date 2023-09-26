package v1

import (
	"api/internal/api/v1/users"
	"api/internal/jwt"
	"api/internal/repos"
	"api/types"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

var handlers map[string]func(gr repos.GlobalRepo, sessionUser *types.User, data []byte) (interface{}, error)
var routesNotRequiringAuthorization types.StringSet

type Request struct {
	Action string `json:"action"`
	Token  string `json:"token"`
	Data   []byte `json:"data"`
}

func getNewToken(gr repos.GlobalRepo, sessionUser *types.User, data []byte) (interface{}, error) {
	return struct {
		Token string `json:"token"`
	}{Token: jwt.NewToken(sessionUser)}, nil
}

func login(gr repos.GlobalRepo, sessionUser *types.User, data []byte) (interface{}, error) {
	type loginRequest struct {
		Email    string `validate:"required" json:"email"`
		Password string `validate:"required" json:"password"`
	}
	req := &loginRequest{}
	if err := json.Unmarshal(data, req); err != nil {
		return nil, err
	}

	if err := types.Validate(req); err != nil {
		return nil, err
	}

	usr, err := gr.Users().GetByEmail(req.Email)
	if err != nil {
		return nil, types.NewNotFoundError("credentials not found")
	}

	if !usr.PasswordMatches(req.Password) {
		return nil, types.NewNotFoundError("credentials not found")
	}

	return struct {
		Token string `json:"token"`
	}{Token: jwt.NewToken(usr)}, nil
}

func init() {
	routesNotRequiringAuthorization = types.NewStringSet("POST:Login")
	handlers = make(map[string]func(repos.GlobalRepo, *types.User, []byte) (interface{}, error))

	handlers["GET:NewToken"] = getNewToken
	handlers["POST:Login"] = login
	handlers = users.HandleRoutes(handlers)
}

func HandleRequest(conn *websocket.Conn, gr repos.GlobalRepo, data []byte) (interface{}, error) {
	req := &Request{}
	if err := json.Unmarshal(data, req); err != nil {
		return nil, fmt.Errorf("unable to read data - must have an action and data fields err: '%s'", err.Error())
	}

	// If route is not in the unauthorized list verify token
	var user *types.User
	if !routesNotRequiringAuthorization.Contains(req.Action) {
		var err error
		user, err = jwt.IsTokenValid(req.Token)
		if err != nil {
			return nil, types.NewUnauthorizedError(err.Error())
		}
	}

	handler, exists := handlers[req.Action]
	if !exists {
		return nil, fmt.Errorf("action '%s' not found", req.Action)
	}

	return handler(gr, user, req.Data)
}
