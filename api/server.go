package server

import (
	"log"
	"net/http"
	"time"

	"gwi/platform2.0-go-challenge/api/modules/audience"
	"gwi/platform2.0-go-challenge/api/modules/authentication"
	"gwi/platform2.0-go-challenge/api/modules/charts"
	"gwi/platform2.0-go-challenge/api/modules/dashboard"
	"gwi/platform2.0-go-challenge/api/modules/health"
	"gwi/platform2.0-go-challenge/api/modules/insights"
	"gwi/platform2.0-go-challenge/environment"

	gwierrors "gwi/platform2.0-go-challenge/pkg/errors"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/gorilla/mux"
)

// ApplicationServer : Instance of an application server.
type ApplicationServer struct {
	HttpServer *http.Server
	Router     http.Handler
	App        *Application
	Config     *environment.Config
}

// NewApplicationServer : Returns an istance of application server.
func NewApplicationServer() *ApplicationServer {
	applicationServer := ApplicationServer{
		Config: environment.LoadConfig(),
	}

	return &applicationServer
}

// Setup : Initiates the application server.
func (o *ApplicationServer) Setup() {
	app := NewApplication(o.Config)

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(o.NotFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(o.MethodNotAllowed)

	health.Setup(router)
	authentication.Setup(router, app.Authentication)
	insights.Setup(router, app.Insights)
	charts.Setup(router, app.Charts)
	audience.Setup(router, app.Audience)
	dashboard.Setup(router, app.Dashboard, app.Charts, app.Audience, app.Insights)

	o.Router = router
	o.HttpServer = &http.Server{
		Addr:              o.Config.ApiAddress,
		Handler:           o.Router,
		ReadTimeout:       10 * time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      60 * time.Second,
	}

	o.App = app
}

// Run : Executes the application server.
func (o *ApplicationServer) Run() {
	log.Println("Server is up and listening ...")
	log.Fatal(http.ListenAndServe(o.HttpServer.Addr, o.Router))
}

// NotFound : Page not found handler.
func (o *ApplicationServer) NotFound(w http.ResponseWriter, r *http.Request) {
	gwihttp.ResponseWithJSON(http.StatusNotFound, gwierrors.NotFoundNew("the page not found"), w)
}

// MethodNotAllowed : Method not allowed handler.
func (o *ApplicationServer) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	gwihttp.ResponseWithJSON(http.StatusMethodNotAllowed, gwierrors.MethodNotAllowedNew("the method is not allowed"), w)
}
