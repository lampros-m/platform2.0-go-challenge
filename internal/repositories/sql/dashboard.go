package sql

import (
	"context"
	"fmt"

	"gwi/platform2.0-go-challenge/internal/app/dashboard"
	"gwi/platform2.0-go-challenge/internal/repositories/tables"

	sq "github.com/Masterminds/squirrel"
)

// Dashboard : Indicates Dashboard repository.
type Dashboard struct {
	client BasicConnectionWithTransactions
}

// NewDashboardRepo : Dashboard repository constructor.
func NewDashboardRepo(client BasicConnectionWithTransactions) *Dashboard {
	return &Dashboard{
		client: client,
	}
}

// ListAssets : List all assets.
func (o *Dashboard) ListAssets(ctx context.Context) (map[uint32]dashboard.Asset, error) {
	w := sq.And{}

	q := sq.Select(
		"a.id",
		"a.title",
		"a.description",
		"a.type",
		"a.created_at",
		"a.updated_at",
	).From(tables.GwiAssets + " AS a").
		Where(w)

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return map[uint32]dashboard.Asset{}, nil
	}
	defer rows.Close()

	out := make(map[uint32]dashboard.Asset)
	for rows.Next() {
		asset := dashboard.Asset{}
		err := rows.Scan(
			&asset.ID,
			&asset.Title,
			&asset.Description,
			&asset.Type,
			&asset.CreatedAt,
			&asset.UpdatedAt,
		)
		if err != nil {
			return map[uint32]dashboard.Asset{}, nil
		}

		out[asset.ID] = asset
	}
	if rows.Err() != nil {
		return map[uint32]dashboard.Asset{}, nil
	}

	return out, nil
}

// GetAssets : Get assets for specific user.
func (o *Dashboard) GetAssets(ctx context.Context, userID uint32) (dashboard.Assets, error) {
	w := sq.And{
		sq.Eq{"ua.user_id": userID},
		sq.Eq{"ua.subscription": true},
	}

	q := sq.Select(
		"a.id",
		"a.title",
		"a.description",
		"a.type",
		"a.created_at",
		"a.updated_at",
	).From(tables.GwiAssets + " AS a").
		InnerJoin(tables.GwiUsersAssets + " AS ua ON a.id = ua.asset_id").
		Where(w)

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return dashboard.Assets{}, nil
	}
	defer rows.Close()

	out := dashboard.Assets{}
	for rows.Next() {
		asset := dashboard.Asset{}
		err := rows.Scan(
			&asset.ID,
			&asset.Title,
			&asset.Description,
			&asset.Type,
			&asset.CreatedAt,
			&asset.UpdatedAt,
		)
		if err != nil {
			return dashboard.Assets{}, nil
		}

		out = append(out, asset)
	}
	if rows.Err() != nil {
		return dashboard.Assets{}, nil
	}

	return out, nil
}

// UpdateAssetDescription : Updates an asset description.
func (o *Dashboard) UpdateAssetDescription(ctx context.Context, description string, assetID dashboard.AssetID) (uint32, error) {
	w := sq.And{
		sq.Eq{"id": assetID},
	}

	q := sq.
		Update(tables.GwiAssets).
		Set("description", description).
		Where(w).
		RunWith(o.client)

	result, err := q.ExecContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("error on sql repo when updating asset description: %s", err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	return uint32(rowsAffected), nil
}

// Subscription : Subsribes or unsubscribes a user to an asset.
func (o *Dashboard) Subscription(ctx context.Context, userID uint32, assetID dashboard.AssetID, subscription bool) error {
	insertBuilder := sq.
		Insert(tables.GwiUsersAssets).
		Columns(
			"user_id",
			"asset_id",
			"subscription",
		).
		Suffix("ON DUPLICATE KEY UPDATE").
		Suffix("subscription = VALUES(subscription)")

	tx, err := o.client.Begin()
	if err != nil {
		return err
	}

	insertBuilder = insertBuilder.Values(
		userID,
		assetID,
		subscription,
	)

	if _, err = insertBuilder.RunWith(tx).Exec(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
