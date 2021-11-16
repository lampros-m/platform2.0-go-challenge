package sql

import (
	"context"

	"gwi/platform2.0-go-challenge/internal/app/insights"
	"gwi/platform2.0-go-challenge/internal/repositories/tables"
	"gwi/platform2.0-go-challenge/pkg/gwitime"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"

	sq "github.com/Masterminds/squirrel"
)

// Insights : Indicates Insights repository.
type Insights struct {
	client BasicConnectionWithTransactions
}

// NewInsightsRepo : Insights repository constructor.
func NewInsightsRepo(client BasicConnectionWithTransactions) *Insights {
	return &Insights{
		client: client,
	}
}

// GetInsights : Retrieves insights.
func (o *Insights) GetInsights(
	ctx context.Context,
	dateFrom gwitime.DateTime,
	dateTo gwitime.DateTime,
	insightType insights.InsightType,
	pgn pagination.PageInfoRequest,
	srt sorting.Sorting) (insights.Insights, error) {
	var direction sorting.Direction

	switch srt.Direction {
	case sorting.DirectionDesc:
		direction = sorting.DirectionDesc
	case sorting.DirectionAsc:
		direction = sorting.DirectionAsc
	default:
		direction = sorting.DirectionDesc
	}

	orderBy := "i.date_produced" + " " + string(direction)

	w := sq.And{
		sq.Eq{"i.type": insightType},
		sq.Expr("DATE(i.date_produced) >= ? AND DATE(i.date_produced) <= ?", dateFrom.Format(gwitime.GWILayoutDateOnly), dateTo.Format(gwitime.GWILayoutDateOnly)),
	}

	q := sq.Select(
		"i.asset_id",
		"i.date_produced",
		"i.message_produced",
		"i.type",
	).
		From(tables.GwiInsights + " AS i").
		InnerJoin(tables.GwiAssets + " AS a ON i.asset_id = a.id").
		Where(w).
		OrderBy(orderBy).
		Offset(uint64(pgn.Offset())).
		Limit(uint64(pgn.PageSize))

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return insights.Insights{}, err
	}
	defer rows.Close()

	out := insights.Insights{}
	for rows.Next() {
		insight := insights.Insight{}
		dateProduced := gwitime.DateTime{}

		err := rows.Scan(
			&insight.AssetID,
			&dateProduced,
			&insight.MessageProduced,
			&insight.Type,
		)
		if err != nil {
			return insights.Insights{}, err
		}

		insight.DateProduced = dateProduced.Format(gwitime.GWILayoutDateOnly)
		out = append(out, insight)
	}
	if rows.Err() != nil {
		return insights.Insights{}, err
	}

	return out, nil
}
