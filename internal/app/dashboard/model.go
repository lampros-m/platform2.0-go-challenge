package dashboard

import (
	"encoding/json"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
)

// AssetID : Describes an asset id.
type AssetID uint32

// AssetType : Represents an asset type.
type AssetType string

// Asset types.
const (
	Audience = AssetType("audience")
	Chart    = AssetType("chart")
	Insight  = AssetType("insight")
)

// Asset : Describes an asset.
type Asset struct {
	ID           uint32           `json:"id"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Type         AssetType        `json:"type"`
	CreatedAt    gwitime.DateTime `json:"created_at"`
	UpdatedAt    gwitime.DateTime `json:"updated_at"`
	Enriched     bool             `json:"enriched"`
	EnrichedInfo interface{}      `json:"enriched_info,omitempty"`
}

// Assets : A slice of assets.
type Assets []Asset

// MarshalJSON : Custom marshal for asset.
func (o *Asset) MarshalJSON() ([]byte, error) {
	type AliasAsset Asset
	return json.Marshal(&struct {
		*AliasAsset
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		AliasAsset: (*AliasAsset)(o),
		CreatedAt:  o.CreatedAt.Format(gwitime.GWILayoutDateOnly),
		UpdatedAt:  o.UpdatedAt.Format(gwitime.GWILayoutDateOnly),
	})
}
