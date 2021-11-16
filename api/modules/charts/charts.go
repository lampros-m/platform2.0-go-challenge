package charts

import (
	"net/http"

	"gwi/platform2.0-go-challenge/api/middleware"
	"gwi/platform2.0-go-challenge/internal/app/charts"
	gwierrors "gwi/platform2.0-go-challenge/pkg/errors"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/gorilla/mux"
)

// Module : Describes a collection of methods for charts module layer.
type Module struct {
	ChartsService *charts.Service
}

// Setup : Setups routes for charts module.
func Setup(router *mux.Router, chartsService *charts.Service) {
	m := &Module{
		ChartsService: chartsService,
	}

	Authenticator := middleware.Authenticator

	charts := router.PathPrefix("/charts").Subrouter()
	charts.HandleFunc("/chartvisits", Authenticator(m.GetChartVisits)).Methods("POST")
	charts.HandleFunc("/chartaudiencereach", Authenticator(m.GetChartAudienceReach)).Methods("POST")
}

// GetChartVisits : Returns a chart for platform visits.
func (o *Module) GetChartVisits(w http.ResponseWriter, r *http.Request) {
	req := GetChartVisitsRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	ctx := r.Context()
	chart, err := o.ChartsService.GetChartVisits(ctx, req.DateFrom, req.DateTo, req.GoogleTraffic)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	gwihttp.ResponseWithJSON(http.StatusOK, chart, w)
}

// GetChartAudienceReach : Returns a chart for audience reach.
func (o *Module) GetChartAudienceReach(w http.ResponseWriter, r *http.Request) {
	req := GetChartAudienceReachRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	ctx := r.Context()
	chart, err := o.ChartsService.GetChartAudienceReach(ctx, req.DateFrom, req.DateTo, req.Reacted)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	gwihttp.ResponseWithJSON(http.StatusOK, chart, w)
}
