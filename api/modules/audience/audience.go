package audience

import (
	"context"
	"errors"
	"net/http"

	"gwi/platform2.0-go-challenge/api/middleware"
	"gwi/platform2.0-go-challenge/internal/app/audience"
	gwierrors "gwi/platform2.0-go-challenge/pkg/errors"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/gorilla/mux"
)

// Module : Describes a collection of methods for audience module layer.
type Module struct {
	AudienceService *audience.Service
}

// Setup : Setups routes for audience module.
func Setup(router *mux.Router, audienceService *audience.Service) {
	m := &Module{
		AudienceService: audienceService,
	}

	Authenticator := middleware.Authenticator

	audience := router.PathPrefix("/audience").Subrouter()
	audience.HandleFunc("/audiencesocialmedia", Authenticator(m.GetAudienceSocialMedia)).Methods("POST")
	audience.HandleFunc("/audienceshopping", Authenticator(m.GetAudienceShopping)).Methods("POST")
}

// GetAudienceSocialMedia : Returns info about audience for social media.
func (o *Module) GetAudienceSocialMedia(w http.ResponseWriter, r *http.Request) {
	req := GetAudienceSocialMediaRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	if req.DateFrom.IsZero() || req.DateTo.IsZero() {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(errors.New("date range is empty")), w)
		return
	}

	req.PageInfoRequest.GetOrDefaultPageInfoRequest(1, 1)

	ctx := r.Context()
	audienceInfo, err := o.AudienceService.GetAudienceSocialMedia(ctx, req.DateFrom, req.DateTo, req.PageInfoRequest, req.Sorting)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	gwihttp.ResponseWithJSON(http.StatusOK, audienceInfo, w)
}

// GetAudienceShopping : Returns info about audience for shopping preferences.
func (o *Module) GetAudienceShopping(w http.ResponseWriter, r *http.Request) {
	req := GetAudienceProductsRequest{}
	if err := gwihttp.ParseAndValidateJSONFromRequest(r, &req); err != nil {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(err), w)
		return
	}

	if req.DateFrom.IsZero() || req.DateTo.IsZero() {
		gwihttp.ResponseWithJSON(http.StatusBadRequest, gwierrors.BadRequest(errors.New("date range is empty")), w)
		return
	}

	req.PageInfoRequest.GetOrDefaultPageInfoRequest(1, 1)

	ctx := r.Context()

	ctx = context.Background()

	audienceInfo, err := o.AudienceService.GetAudienceShopping(ctx, req.DateFrom, req.DateTo, req.PageInfoRequest, req.Sorting)
	if err != nil {
		gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServer(err), w)
		return
	}

	gwihttp.ResponseWithJSON(http.StatusOK, audienceInfo, w)
}
