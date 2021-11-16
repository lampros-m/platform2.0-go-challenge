package sql

import (
	"context"

	"gwi/platform2.0-go-challenge/internal/app/audience"
	"gwi/platform2.0-go-challenge/internal/repositories/tables"
	"gwi/platform2.0-go-challenge/pkg/gwitime"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"

	sq "github.com/Masterminds/squirrel"
)

// Audience : Indicates Audience repository.
type Audience struct {
	client BasicConnectionWithTransactions
}

// NewAudienceRepo : Audience repository constructor.
func NewAudienceRepo(client BasicConnectionWithTransactions) *Audience {
	return &Audience{
		client: client,
	}
}

// GetAudienceSocialMedia : Gets audience info for social media.
func (o *Audience) GetAudienceSocialMedia(
	ctx context.Context,
	dateFrom gwitime.DateTime,
	dateTo gwitime.DateTime,
	pgn pagination.PageInfoRequest,
	srt sorting.Sorting) (audience.AudienceSocialMediaMultiple, error) {
	var direction sorting.Direction

	switch srt.Direction {
	case sorting.DirectionDesc:
		direction = sorting.DirectionDesc
	case sorting.DirectionAsc:
		direction = sorting.DirectionAsc
	default:
		direction = sorting.DirectionDesc
	}

	orderBy := "au.date_produced" + " " + string(direction)

	w := sq.And{
		sq.Expr("DATE(au.date_produced) >= ? AND DATE(au.date_produced) <= ?", dateFrom.Format(gwitime.GWILayoutDateOnly), dateTo.Format(gwitime.GWILayoutDateOnly)),
	}

	q := sq.Select(
		"au.asset_id",
		"au.date_produced",
		"au.gender",
		"c.country_name",
		"au.age_from",
		"au.age_to",
		"au.hours_spent",
		"s.media_name",
	).From(tables.GwiAudienceSocialMedia + " AS au").
		InnerJoin(tables.GwiAssets + " AS a ON au.asset_id = a.id").
		InnerJoin(tables.GwiCountries + " AS c ON au.country = c.id").
		InnerJoin(tables.GwiSocialMedia + " AS s ON au.social_media = s.id").
		Where(w).
		OrderBy(orderBy).
		Offset(uint64(pgn.Offset())).
		Limit(uint64(pgn.PageSize))

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return audience.AudienceSocialMediaMultiple{}, err
	}
	defer rows.Close()

	out := audience.AudienceSocialMediaMultiple{}
	for rows.Next() {
		a := audience.AudienceSocialMedia{}
		date := gwitime.DateTime{}

		err := rows.Scan(
			&a.AssetID,
			&date,
			&a.Gender,
			&a.Country,
			&a.AgeFrom,
			&a.AgeTo,
			&a.HoursSpent,
			&a.SocialMedia,
		)
		if err != nil {
			return audience.AudienceSocialMediaMultiple{}, err
		}

		a.DateProduced = date.Format(gwitime.GWILayoutDateOnly)
		out = append(out, a)
	}
	if rows.Err() != nil {
		return audience.AudienceSocialMediaMultiple{}, err
	}

	return out, nil
}

// GetAudienceShopping : Gets audio info for shopping activities.
func (o *Audience) GetAudienceShopping(
	ctx context.Context,
	dateFrom gwitime.DateTime,
	dateTo gwitime.DateTime,
	pgn pagination.PageInfoRequest,
	srt sorting.Sorting) (audience.AudienceProductsMultiple, error) {

	var direction sorting.Direction

	switch srt.Direction {
	case sorting.DirectionDesc:
		direction = sorting.DirectionDesc
	case sorting.DirectionAsc:
		direction = sorting.DirectionAsc
	default:
		direction = sorting.DirectionDesc
	}

	orderBy := "au.date_produced" + " " + string(direction)

	w := sq.And{
		sq.Expr("DATE(au.date_produced) >= ? AND DATE(au.date_produced) <= ?", dateFrom.Format(gwitime.GWILayoutDateOnly), dateTo.Format(gwitime.GWILayoutDateOnly)),
	}

	q := sq.Select(
		"au.asset_id",
		"au.date_produced",
		"au.gender",
		"c.country_name",
		"au.age_from",
		"au.age_to",
		"au.hours_spent",
		"p.product_name",
	).From(tables.GwiAudienceShopping + " AS au").
		InnerJoin(tables.GwiAssets + " AS a ON au.asset_id = a.id").
		InnerJoin(tables.GwiCountries + " AS c ON au.country = c.id").
		InnerJoin(tables.GwiProducts + " AS p ON au.product = p.id").
		Where(w).
		OrderBy(orderBy).
		Offset(uint64(pgn.Offset())).
		Limit(uint64(pgn.PageSize))

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return audience.AudienceProductsMultiple{}, err
	}
	defer rows.Close()

	out := audience.AudienceProductsMultiple{}
	for rows.Next() {
		a := audience.AudienceProducts{}
		date := gwitime.DateTime{}
		var assetID uint32

		err := rows.Scan(
			&assetID,
			&date,
			&a.Gender,
			&a.Country,
			&a.AgeFrom,
			&a.AgeTo,
			&a.HoursSpent,
			&a.Product,
		)
		if err != nil {
			return audience.AudienceProductsMultiple{}, err
		}

		a.DateProduced = date.Format(gwitime.GWILayoutDateOnly)
		out = append(out, a)
	}
	if rows.Err() != nil {
		return audience.AudienceProductsMultiple{}, err
	}

	return out, nil
}
