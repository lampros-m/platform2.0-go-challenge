package authentication

import (
	"fmt"
	"net/http"

	"gwi/platform2.0-go-challenge/internal/app/authentication"
	"gwi/platform2.0-go-challenge/pkg/broadcast"
	gwierrors "gwi/platform2.0-go-challenge/pkg/errors"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/gorilla/mux"
)

// Module : Describes a collection of methods for authentication module layer.
type Module struct {
	Authentication *authentication.Service
}

// Setup : Setups routs for authentication module.
func Setup(router *mux.Router, authenticationService *authentication.Service) {
	m := &Module{
		Authentication: authenticationService,
	}

	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", m.Signup)
	auth.HandleFunc("/login", m.Login)
}

// Signup : Signup for user.
func (o *Module) Signup(w http.ResponseWriter, r *http.Request) {
	req := SignupLoginRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	ctx := r.Context()
	err := o.Authentication.Signup(ctx, req.Username, req.Password)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	broadcast := broadcast.Broadcast{Message: "sucess"}
	gwihttp.ResponseWithJSON(http.StatusOK, broadcast, w)
}

// Login : Login for user.
func (o *Module) Login(w http.ResponseWriter, r *http.Request) {
	req := SignupLoginRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	user := req.Username
	pass := req.Password

	ctx := r.Context()
	token, err := o.Authentication.Login(ctx, user, pass)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	broadcast := broadcast.Broadcast{
		Message:     fmt.Sprintf("Welcome %s", user),
		Information: []interface{}{map[string]string{"token": token}},
	}
	gwihttp.ResponseWithJSON(http.StatusOK, broadcast, w)
}
