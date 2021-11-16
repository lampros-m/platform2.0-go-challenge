package dashboard

import (
	"errors"
	"fmt"
	"net/http"

	"gwi/platform2.0-go-challenge/api/middleware"
	"gwi/platform2.0-go-challenge/internal/app/audience"
	"gwi/platform2.0-go-challenge/internal/app/charts"
	"gwi/platform2.0-go-challenge/internal/app/dashboard"
	"gwi/platform2.0-go-challenge/internal/app/insights"
	"gwi/platform2.0-go-challenge/pkg/broadcast"
	gwierrors "gwi/platform2.0-go-challenge/pkg/errors"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/gorilla/mux"
)

// Module : Describes a collection of methods for dashboard module layer.
type Module struct {
	DashboardService *dashboard.Service
	ChartsService    *charts.Service
	AudienceService  *audience.Service
	InsightsService  *insights.Service
}

// Setup : Setups routes for dashboard module.
func Setup(
	router *mux.Router,
	dashboardService *dashboard.Service,
	chartsService *charts.Service,
	audienceService *audience.Service,
	insightsService *insights.Service) {
	m := &Module{
		DashboardService: dashboardService,
		ChartsService:    chartsService,
		AudienceService:  audienceService,
		InsightsService:  insightsService,
	}

	Authenticator := middleware.Authenticator

	dashboard := router.PathPrefix("/dashboard").Subrouter()
	dashboard.HandleFunc("/listassets", Authenticator(m.ListAssets)).Methods("GET")
	dashboard.HandleFunc("/userassets", Authenticator(m.UserAssets)).Methods("POST")
	dashboard.HandleFunc("/updateassetdescription", Authenticator(m.UpdateAssetDescription)).Methods("PATCH")
	dashboard.HandleFunc("/subscription", Authenticator(m.Subscription)).Methods("POST")
}

// UserAssets : Returns an asset based on users id.
func (o *Module) UserAssets(w http.ResponseWriter, r *http.Request) {
	req := UserAssetsRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	userID, err := gwihttp.ExtractUserIDFromRequest(r)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusUnauthorized, gwierrors.Unauthorized(err), w)
		return
	}

	ctx := r.Context()
	userassets, err := o.DashboardService.GetAssets(ctx, userID)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	if !req.EnricheView {
		gwihttp.ResponseWithJSON(http.StatusOK, userassets, w)
		return
	}

	o.enricheAssets(w, r, userassets)

	gwihttp.ResponseWithJSON(http.StatusOK, userassets, w)
}

// UpdateAssetDescription : Updates asset's description.
func (o *Module) UpdateAssetDescription(w http.ResponseWriter, r *http.Request) {
	req := UpdateAssetDescriptionRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	if req.ID == 0 {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(errors.New("asset id must be provided")), w)
		return
	}

	ctx := r.Context()
	err := o.DashboardService.UpdateAssetDescription(ctx, req.Description, req.ID)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	broadcast := broadcast.Broadcast{Message: fmt.Sprintf("description asset with id %d updated successfully", req.ID)}
	gwihttp.ResponseWithJSON(http.StatusOK, broadcast, w)
}

// ListAssets : Lists all assets.
func (o *Module) ListAssets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	assets, err := o.DashboardService.ListAssets(ctx)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	gwihttp.ResponseWithJSON(http.StatusOK, assets, w)
}

// Subscription : Subscribe / unsubscribe from asset.
func (o *Module) Subscription(w http.ResponseWriter, r *http.Request) {
	req := SubscriptionRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	userID, err := gwihttp.ExtractUserIDFromRequest(r)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusUnauthorized, gwierrors.Unauthorized(err), w)
		return
	}

	ctx := r.Context()
	err = o.DashboardService.Subscription(ctx, userID, req.ID, req.Subscription)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	broadcastMessage := broadcast.Broadcast{Message: "subscribed successfully"}
	if !req.Subscription {
		broadcastMessage = broadcast.Broadcast{Message: "unsubscribed successfully"}
	}

	gwihttp.ResponseWithJSON(http.StatusOK, broadcastMessage, w)
}
