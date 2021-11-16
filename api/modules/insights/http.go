package insights

import (
	"gwi/platform2.0-go-challenge/internal/app/insights"
	"gwi/platform2.0-go-challenge/pkg/daterange"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"
)

// GetInsightsRequest : Describes a get insight request.
type GetInsightsRequest struct {
	daterange.DateRange
	pagination.PageInfoRequest
	sorting.Sorting
	InsightType insights.InsightType `json:"insight_type"`
}
