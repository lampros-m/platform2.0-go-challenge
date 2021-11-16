package dashboard

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"gwi/platform2.0-go-challenge/internal/app/dashboard"
	"gwi/platform2.0-go-challenge/internal/app/insights"
	"gwi/platform2.0-go-challenge/pkg/daterange"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"
)

func registeredEnrichedAssets(o *Module) map[uint32]func(http.ResponseWriter, *http.Request) (interface{}, error) {
	return map[uint32]func(http.ResponseWriter, *http.Request) (interface{}, error){
		1: o.enrichedAudienceShopping,
		2: o.enrichedAudienceSocialMedia,
		3: o.enrichedChartAudienceReach,
		4: o.enrichedChartVisits,
		5: o.enrichedInsights,
	}
}

func (o *Module) enrichedAudienceShopping(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	defaultDateRange := daterange.GetDefault()
	pgn := pagination.PageInfoRequest{}
	pgn.GetOrDefaultPageInfoRequest(1, 1)
	srt := sorting.DefaultSorting()

	audienceInfo, err := o.AudienceService.GetAudienceShopping(ctx, defaultDateRange.DateFrom, defaultDateRange.DateTo, pgn, srt)
	if err != nil {
		log.Printf("error on enriched audience shopping: %s", err.Error())
		return fmt.Errorf("error on enriched audience shopping: %s", err.Error()), nil
	}

	return audienceInfo, nil
}

func (o *Module) enrichedAudienceSocialMedia(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	defaultDateRange := daterange.GetDefault()
	pgn := pagination.PageInfoRequest{}
	pgn.GetOrDefaultPageInfoRequest(1, 1)
	srt := sorting.DefaultSorting()

	audienceInfo, err := o.AudienceService.GetAudienceSocialMedia(ctx, defaultDateRange.DateFrom, defaultDateRange.DateTo, pgn, srt)
	if err != nil {
		log.Printf("error on enriched audience social media: %s", err.Error())
		return fmt.Errorf("error on enriched social media: %s", err.Error()), nil
	}

	return audienceInfo, nil
}

func (o *Module) enrichedChartAudienceReach(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	defaultDateRange := daterange.GetDefault()
	reacted := false

	chart, err := o.ChartsService.GetChartAudienceReach(ctx, defaultDateRange.DateFrom, defaultDateRange.DateTo, reacted)
	if err != nil {
		log.Printf("error on enriched chart audience reach: %s", err.Error())
		return fmt.Errorf("error on enriched chart audience reach: %s", err.Error()), nil
	}

	return chart, nil
}

func (o *Module) enrichedChartVisits(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	defaultDateRange := daterange.GetDefault()
	reacted := false

	chart, err := o.ChartsService.GetChartVisits(ctx, defaultDateRange.DateFrom, defaultDateRange.DateTo, reacted)
	if err != nil {
		log.Printf("error on enriched chart visits: %s", err.Error())
		return fmt.Errorf("error on enriched chart visits: %s", err.Error()), nil
	}

	return chart, nil
}

func (o *Module) enrichedInsights(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	defaultDateRange := daterange.GetDefault()
	pgn := pagination.PageInfoRequest{}
	pgn.GetOrDefaultPageInfoRequest(1, 1)
	srt := sorting.DefaultSorting()
	insightType := insights.TypeActivity

	insights, err := o.InsightsService.GetInsights(ctx, defaultDateRange.DateFrom, defaultDateRange.DateTo, insightType, pgn, srt)
	if err != nil {
		log.Printf("error on enriched insights: %s", err.Error())
		return fmt.Errorf("error on enriched insights: %s", err.Error()), nil
	}

	return insights, nil
}

func (o *Module) enricheAssets(w http.ResponseWriter, r *http.Request, assets dashboard.Assets) {
	enrichedAssetCh := make(chan dashboard.Asset)
	enricheFuncs := registeredEnrichedAssets(o)
	var wg sync.WaitGroup

	for i := range assets {
		if enricheFunction, found := enricheFuncs[assets[i].ID]; found {
			wg.Add(1)
			go func(f func(http.ResponseWriter, *http.Request) (interface{}, error), a dashboard.Asset) {
				defer wg.Done()

				specialInfo, err := enricheFunction(w, r)
				if err != nil {
					log.Printf("error on fetching special info: %s", err.Error())
					enrichedAssetCh <- a
					return
				}

				a.EnrichedInfo = specialInfo
				if specialInfo != nil {
					a.Enriched = true
				}

				enrichedAssetCh <- a
			}(enricheFunction, assets[i])
			continue
		}

		wg.Add(1)
		go func(a dashboard.Asset) {
			defer wg.Done()
			enrichedAssetCh <- a
		}(assets[i])
	}

	go func() {
		wg.Wait()
		close(enrichedAssetCh)
	}()

	enrichedAssets := make(dashboard.Assets, 0, len(assets))
	for ea := range enrichedAssetCh {
		enrichedAssets = append(enrichedAssets, ea)
	}

	copy(assets, enrichedAssets)
}
