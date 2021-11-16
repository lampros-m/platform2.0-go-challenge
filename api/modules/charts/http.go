package charts

import "gwi/platform2.0-go-challenge/pkg/daterange"

// GetChartVisitsRequest : A request for chart visits.
type GetChartVisitsRequest struct {
	daterange.DateRange
	GoogleTraffic bool `json:"google_traffic"`
}

// GetChartAudienceReachRequest : A request for audience reach chart.
type GetChartAudienceReachRequest struct {
	daterange.DateRange
	Reacted bool `json:"reacted"`
}
