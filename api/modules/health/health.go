package health

import (
	"net/http"

	"gwi/platform2.0-go-challenge/pkg/broadcast"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/gorilla/mux"
)

// Module : A health module.
type Module struct{}

// Setup : Setups routes from health module.
func Setup(router *mux.Router) {
	router.HandleFunc("/health", Health).Methods("GET", "POST")
}

// Helath : Checks API health.
func Health(w http.ResponseWriter, r *http.Request) {
	message := broadcast.Broadcast{
		Message: "Alive!",
	}

	gwihttp.ResponseWithJSON(http.StatusOK, message, w)
}
