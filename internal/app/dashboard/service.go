package dashboard

import (
	"context"
	"fmt"
	"log"
)

// Service : Struct that represents dashboard service.
type Service struct {
	repo Repository
}

// NewService : Service dashboard constructor.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetAssets : Returns assets.
func (o *Service) GetAssets(ctx context.Context, userID uint32) (Assets, error) {
	return o.repo.GetAssets(ctx, userID)
}

// ListAssets : Lists all assets.
func (o *Service) ListAssets(ctx context.Context) (Assets, error) {
	assetsMap, err := o.repo.ListAssets(ctx)
	if err != nil {
		return Assets{}, fmt.Errorf("error on fetching list of assets: %s", err.Error())
	}

	assets := make(Assets, 0, len(assetsMap))
	for _, a := range assetsMap {
		assets = append(assets, a)
	}

	return assets, nil
}

// UpdateAssetDescription : Updates asset's description.
func (o *Service) UpdateAssetDescription(ctx context.Context, description string, assetID AssetID) error {
	rowsAffected, err := o.repo.UpdateAssetDescription(ctx, description, assetID)
	if err != nil {
		log.Printf("error on asset subscription: %s", err.Error())
		return fmt.Errorf("error on asset subscription: %s", err.Error())
	}

	if rowsAffected > 1 {
		log.Printf("error on updating asset description. more than 2 rows updated for asset: %d", assetID)
		return fmt.Errorf("error on updating asset description. more than 2 rows updated for asset: %d", assetID)
	}

	return nil
}

// Subscription : Subscribes and unsubsribes a user from an asset.
func (o *Service) Subscription(ctx context.Context, userID uint32, assetID AssetID, subscription bool) error {
	assetsMap, err := o.repo.ListAssets(ctx)
	if err != nil {
		return fmt.Errorf("error on fetching list of assets: %s", err.Error())
	}

	if _, found := assetsMap[uint32(assetID)]; !found {
		return fmt.Errorf("error - cannot find asset: %d", assetID)
	}

	return o.repo.Subscription(ctx, userID, assetID, subscription)
}
