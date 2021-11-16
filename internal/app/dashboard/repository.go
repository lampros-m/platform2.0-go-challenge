package dashboard

import "context"

// Repository : Interface for dashboard repository.
type Repository interface {
	ListAssets(ctx context.Context) (map[uint32]Asset, error)
	UpdateAssetDescription(ctx context.Context, description string, assetID AssetID) (uint32, error)
	Subscription(ctx context.Context, userID uint32, assetID AssetID, subscription bool) error
	GetAssets(ctx context.Context, userID uint32) (Assets, error)
}
