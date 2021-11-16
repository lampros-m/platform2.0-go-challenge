package insights

import (
	"context"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"
)

// Repository : Interface for insights repository.
type Repository interface {
	GetInsights(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, insightType InsightType, pagination pagination.PageInfoRequest, sorting sorting.Sorting) (Insights, error)
}
