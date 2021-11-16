package audience

import (
	"context"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"
)

// Repository : Interface for audience repository.
type Repository interface {
	GetAudienceSocialMedia(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, pgn pagination.PageInfoRequest, srt sorting.Sorting) (AudienceSocialMediaMultiple, error)
	GetAudienceShopping(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, pgn pagination.PageInfoRequest, srt sorting.Sorting) (AudienceProductsMultiple, error)
}
