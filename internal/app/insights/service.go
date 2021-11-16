package insights

import (
	"context"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"
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

// GetInsights : Returns insights.
func (o *Service) GetInsights(
	ctx context.Context,
	dateFrom gwitime.DateTime,
	dateTo gwitime.DateTime,
	insightType InsightType,
	pagination pagination.PageInfoRequest,
	sorting sorting.Sorting) (Insights, error) {
	return o.repo.GetInsights(ctx, dateFrom, dateTo, insightType, pagination, sorting)
}
