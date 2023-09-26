package apipkg

import (
	v1 "api/internal/api/v1"
	"api/internal/repos"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

			res, herr := v1.HandleRequest(conn, gr, p)
			bts, err := json.Marshal(struct {
				Data  interface{} `json:"data"`
				Error error       `json:"error"`
			}{
				Data:  res,
				Error: herr,
			})
			if err != nil {
				log.Println(err)
				return
			}

			if err := conn.WriteMessage(1, bts); err != nil {
				log.Println(err)
				return
			}
		}
	})

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting server on port '%d'\n", port)
	log.Fatal(srv.ListenAndServe())
}
