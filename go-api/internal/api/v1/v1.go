package v1

import (
	socketHandlerUsers "api/internal/api/v1/socket/users"
	"api/internal/jwt"
	"api/internal/repos"
	"api/internal/utils"
	"api/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var handlers map[string]func(gr repos.GlobalRepo, sessionUser *types.User, data []byte) *types.WSResponse
var routesNotRequiringAuthorization types.StringSet

func GetNewToken(gr repos.GlobalRepo, sessionUser *types.User, data []byte) *types.WSResponse {
	return utils.ReqSuccess(struct {
		Token string `json:"token"`
	}{Token: jwt.NewToken(sessionUser)})
}

func Login(gr repos.GlobalRepo, sessionUser *types.User, data []byte) *types.WSResponse {
	type loginRequest struct {
		Email    string `validate:"required" json:"email"`
		Password string `validate:"required" json:"password"`
	}

	req := &loginRequest{}
	if reqErr := utils.ParseSocketReq(data, req); reqErr != nil {
		return utils.ReqFailBadRequest("email and password required to login")
	}

	usr, err := gr.Users().GetByEmail(req.Email)
	if err != nil {
		return utils.ReqFailForbidden("credentials not found")
	}

	if !usr.PasswordMatches(req.Password) {
		return utils.ReqFailForbidden("credentials not found")
	}

	return utils.ReqSuccess(struct {
		Token string `json:"token"`
	}{Token: jwt.NewToken(usr)})
}

func init() {
	routesNotRequiringAuthorization = types.NewStringSet("POST:Login")
	handlers = make(map[string]func(repos.GlobalRepo, *types.User, []byte) *types.WSResponse)

	handlers["GET:NewToken"] = GetNewToken
	handlers["POST:Login"] = Login
	handlers = socketHandlerUsers.HandleRoutes(handlers)
}

func HandleSocketRequest(conn *websocket.Conn, gr repos.GlobalRepo, data []byte) *types.WSResponse {
	req := &types.WSRequest{}
	if err := json.Unmarshal(data, req); err != nil {
		return &types.WSResponse{
			Status: http.StatusBadRequest,
			Error:  "unable to read data - must have an action and data fields",
		}
	}

	if err := types.Validate(req); err != nil {
		return &types.WSResponse{
			Status: http.StatusBadRequest,
			Error:  "unable to read data - must have an action and data fields",
		}
	}

	// If route is not in the unauthorized list verify token
	var user *types.User
	if !routesNotRequiringAuthorization.Contains(req.Action) {
		var err error
		user, err = jwt.IsTokenValid(req.Token)
		if err != nil {
			return &types.WSResponse{
				ID:     req.ID,
				Status: http.StatusForbidden,
				Error:  types.NewUnauthorizedError(err.Error()).Error(),
			}
		}
	}

	handler, exists := handlers[req.Action]
	if !exists {
		return &types.WSResponse{
			ID:     req.ID,
			Status: http.StatusForbidden,
			Error:  fmt.Errorf("action '%s' not found", req.Action).Error(),
		}
	}

	res := handler(gr, user, req.Data)
	res.ID = req.ID
	return res
}

func HandleHTTPRequests(gr repos.GlobalRepo, subrouter *mux.Router) {
	subrouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("X-App-Token")
			if len(token) < 1 {
				http.Error(w, "Not authorized", http.StatusUnauthorized)
				return
			}

			_, err := jwt.IsTokenValid(token)
			if err != nil {
				http.Error(w, "Not authorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.Background()))
		})
	})

	subrouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})
}
