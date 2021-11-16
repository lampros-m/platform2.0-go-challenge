package charts

import (
	"context"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
)

// Repository : Interface for charts repository.
type Repository interface {
	GetChartVisits(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, googleTraffic bool) (VisitsChart, error)
	GetChartAudienceReach(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, reacted bool) (AudienceReachChart, error)
}
