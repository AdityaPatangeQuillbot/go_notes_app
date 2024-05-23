package core

import (
	"encoding/json"
	"log"
	"net/http"
)

func FormatErrorResponseJSON(w *http.ResponseWriter, err error, status int) error {
	response := map[string]interface{}{
		"error":       err.Error(),
		"status_code": status,
	}

	(*w).WriteHeader(status)
	err_ := json.NewEncoder(*w).Encode(response)

	if err_ != nil {
		log.Fatal("error encoding response to json", err)
		return err
	}

	return nil
}
