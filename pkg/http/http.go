package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"gwi/platform2.0-go-challenge/environment"
)

// ResponseWithJSON : Wraps the response process.
func ResponseWithJSON(status int, i interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(i)
}

// ParseAndValidateJSONFromRequest function.
func ParseAndValidateJSONFromRequest(r *http.Request, i interface{}) error {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return nil
}

// ExtractUserIDFromRequest : Extracts gwi user id from request.
func ExtractUserIDFromRequest(r *http.Request) (uint32, error) {
	config := environment.LoadConfig()
	userIndicaton := config.GwiUser

	userID, err := strconv.Atoi(r.Header.Get(userIndicaton))
	if err != nil {
		return 0, err
	}

	userIDformated := uint32(userID)
	if userIDformated == 0 {
		return 0, errors.New("user id cannot be 0")
	}

	return uint32(userID), nil
}
