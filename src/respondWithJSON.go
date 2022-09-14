package src

import (
	"encoding/json"
	"net/http"
)

// Responding with JSON format
func RespondWithJSON(w http.ResponseWriter, status int, resp interface{}) error {
	response, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
	return nil
}
