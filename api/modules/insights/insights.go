package insights

import (
	"net/http"

	"gwi/platform2.0-go-challenge/api/middleware"
	"gwi/platform2.0-go-challenge/internal/app/insights"
	gwierrors "gwi/platform2.0-go-challenge/pkg/errors"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/gorilla/mux"
)

// Module : Describes a collection of methods for insight module layer.
type Module struct {
	IngightsService *insights.Service
}

// Setup : Setups routes for insights module.
func Setup(router *mux.Router, insightsService *insights.Service) {
	m := &Module{
		IngightsService: insightsService,
	}

	Authenticator := middleware.Authenticator

	insights := router.PathPrefix("/insights").Subrouter()
	insights.HandleFunc("/insights", Authenticator(m.GetInsights)).Methods("POST")
}

// GetInsights : Retrives insights messages.
func (o *Module) GetInsights(w http.ResponseWriter, r *http.Request) {
	req := GetInsightsRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	req.PageInfoRequest.GetOrDefaultPageInfoRequest(1, 1)

	ctx := r.Context()
	insights, err := o.IngightsService.GetInsights(ctx, req.DateFrom, req.DateTo, req.InsightType, req.PageInfoRequest, req.Sorting)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	// TODO : Return page info.

	gwihttp.ResponseWithJSON(http.StatusOK, insights, w)
}
