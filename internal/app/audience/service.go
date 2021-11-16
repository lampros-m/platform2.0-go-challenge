package audience

import (
	"context"
	"fmt"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"
)

// Service : Struct that represents audience service.
type Service struct {
	repo Repository
}

// NewService : Service audience constructor.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetAudienceSocialMedia : Audience information about social meda.
func (o *Service) GetAudienceSocialMedia(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, pgn pagination.PageInfoRequest, srt sorting.Sorting) ([]string, error) {
	audienceInfo, err := o.repo.GetAudienceSocialMedia(ctx, dateFrom, dateTo, pgn, srt)
	if err != nil {
		return []string{}, fmt.Errorf("error on fetching info about social media audience: %s", err)
	}

	messagesProduced := make([]string, 0, len(audienceInfo))
	for _, a := range audienceInfo {
		messagesProduced = append(messagesProduced, a.ProduceMessage())
	}

	return messagesProduced, nil
}

// GetAudienceShopping : Audience information about shopping activites.
func (o *Service) GetAudienceShopping(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, pgn pagination.PageInfoRequest, srt sorting.Sorting) ([]string, error) {
	audienceInfo, err := o.repo.GetAudienceShopping(ctx, dateFrom, dateTo, pgn, srt)
	if err != nil {
		return []string{}, fmt.Errorf("error on fetching info about shopping audience: %s", err)
	}

	messagesProduced := make([]string, 0, len(audienceInfo))
	for _, a := range audienceInfo {
		messagesProduced = append(messagesProduced, a.ProduceMessage())
	}

	return messagesProduced, nil
}
