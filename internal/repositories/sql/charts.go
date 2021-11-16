package sql

import (
	"context"

	"gwi/platform2.0-go-challenge/internal/app/charts"
	"gwi/platform2.0-go-challenge/internal/repositories/tables"
	"gwi/platform2.0-go-challenge/pkg/gwitime"

	sq "github.com/Masterminds/squirrel"
)

// Charts : Indicates Charts repository.
type Charts struct {
	client BasicConnectionWithTransactions
}

// NewChartsRepo : Charts repository constructor.
func NewChartsRepo(client BasicConnectionWithTransactions) *Charts {
	return &Charts{
		client: client,
	}
}

// GetChartVisits : Get chart for visits.
func (o *Charts) GetChartVisits(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, googleTraffic bool) (charts.VisitsChart, error) {
	w := sq.And{
		sq.Eq{"vi.google_traffic": googleTraffic},
		sq.Expr("DATE(vi.date) >= ? AND DATE(vi.date) <= ?", dateFrom.Format(gwitime.GWILayoutDateOnly), dateTo.Format(gwitime.GWILayoutDateOnly)),
	}

	q := sq.Select(
		"vi.asset_id",
		"vi.date",
		"vi.visits",
	).From(tables.GwiChartVisits + " AS vi").
		InnerJoin(tables.GwiAssets + " AS a ON vi.asset_id = a.id").
		Where(w)

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return charts.VisitsChart{}, err
	}
	defer rows.Close()

	visitsPeriod := []charts.VisitsPerPeriod{}
	var assetID uint32
	for rows.Next() {
		visits := charts.VisitsPerPeriod{}
		date := gwitime.DateTime{}

		err := rows.Scan(
			&assetID,
			&date,
			&visits.Visits,
		)
		if err != nil {
			return charts.VisitsChart{}, err
		}

		visits.Period = date.Format(gwitime.GWILayoutDateOnly)
		visitsPeriod = append(visitsPeriod, visits)
	}
	if rows.Err() != nil {
		return charts.VisitsChart{}, err
	}

	out := charts.VisitsChart{
		VisitsPerPeriod: visitsPeriod,
		AssetID:         assetID,
	}

	return out, nil
}

// GetChartAudienceReach : Get chart for audience reach.
func (o *Charts) GetChartAudienceReach(ctx context.Context, dateFrom gwitime.DateTime, dateTo gwitime.DateTime, reacted bool) (charts.AudienceReachChart, error) {
	w := sq.And{
		sq.Eq{"au.audience_reacted": reacted},
		sq.Expr("DATE(au.date) >= ? AND DATE(au.date) <= ?", dateFrom.Format(gwitime.GWILayoutDateOnly), dateTo.Format(gwitime.GWILayoutDateOnly)),
	}

	q := sq.Select(
		"au.asset_id",
		"au.date",
		"au.audience_count",
	).From(tables.GwiChartAudienceReach + " AS au").
		InnerJoin(tables.GwiAssets + " AS a ON au.asset_id = a.id").
		Where(w)

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return charts.AudienceReachChart{}, err
	}
	defer rows.Close()

	audienceReache := []charts.AudienceReachPerPeriod{}
	var assetID uint32

	for rows.Next() {
		reach := charts.AudienceReachPerPeriod{}
		date := gwitime.DateTime{}

		err := rows.Scan(
			&assetID,
			&date,
			&reach.Reached,
		)
		if err != nil {
			return charts.AudienceReachChart{}, err
		}

		reach.Period = date.Format(gwitime.GWILayoutDateOnly)
		audienceReache = append(audienceReache, reach)
	}
	if rows.Err() != nil {
		return charts.AudienceReachChart{}, err
	}

	out := charts.AudienceReachChart{
		AudienceReachPerPeriod: audienceReache,
		AssetID:                assetID,
	}

	return out, nil
}
