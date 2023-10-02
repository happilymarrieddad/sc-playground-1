package apipkg

import (
	v1 "api/internal/api/v1"
	"api/internal/repos"
	"api/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Start(gr repos.GlobalRepo, port int) {
	r := mux.NewRouter()

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			bts, _ := json.Marshal(v1.HandleSocketRequest(conn, gr, p))
			if err := conn.WriteMessage(1, bts); err != nil {
				log.Println(err)
				return
			}
		}
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		body, err := utils.HandleHTTPRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := v1.Login(gr, nil, body)
		if len(res.Error) > 0 {
			http.Error(w, res.Error, res.Status)
			return
		}

		utils.HandleHTTPResponse(w, res.Data)
	}).Methods("POST")

	v1.HandleHTTPRequests(gr, r.PathPrefix("/v1").Subrouter())

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	srv := &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting server on port '%d'\n", port)
	log.Fatal(srv.ListenAndServe())
}
