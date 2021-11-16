package dashboard

import "gwi/platform2.0-go-challenge/internal/app/dashboard"

// UpdateAssetDescriptionRequest : Describes an update asset description request.
type UpdateAssetDescriptionRequest struct {
	ID          dashboard.AssetID `json:"id"`
	Description string            `json:"description"`
}

// SubscriptionRequest : Describes subscribe or unsubsribe request.
type SubscriptionRequest struct {
	ID           dashboard.AssetID `json:"id"`
	Subscription bool              `json:"subscription"`
}

// UserAssetsRequest : Describes a user assets request.
type UserAssetsRequest struct {
	EnricheView bool `json:"enriched_view"`
}
