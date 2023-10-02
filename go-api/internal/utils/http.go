package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func HandleHTTPRequest(w http.ResponseWriter, r *http.Request) (body json.RawMessage, err error) {
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, errors.New("unable to parse request")
	}

	return body, nil
}

func HandleHTTPResponse(w http.ResponseWriter, data interface{}) {
	bts, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "unable to parse response", http.StatusInternalServerError)
		return
	}

	w.Write(bts)
}
