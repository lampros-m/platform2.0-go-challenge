package charts

import (
	"context"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
)

// Service : Struct that represents charts service.
type Service struct {
	repo Repository
}

// NewService : Service charts constructor.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetVisits : Returns visits chart.
func (o *Service) GetChartVisits(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, googleTraffic bool) (VisitsChart, error) {
	return o.repo.GetChartVisits(ctx, dateFrom, dateTo, googleTraffic)
}

// GetChartAudienceReach : Returns audience reach chart.
func (o *Service) GetChartAudienceReach(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, reacted bool) (AudienceReachChart, error) {
	return o.repo.GetChartAudienceReach(ctx, dateFrom, dateTo, reacted)
}
