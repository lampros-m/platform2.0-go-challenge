package insights

import (
	"fmt"
	"strings"
)

// Insight : Describes info about insight.
type Insight struct {
	AssetID         uint32         `json:"-"`
	DateProduced    string         `json:"date_produced"`
	MessageProduced InsightMessage `json:"message"`
	Type            InsightType    `json:"type"`
}

// Insights : Describes a slice of insights.
type Insights []Insight

// InsightMessage : Describes an insight message.
type InsightMessage string

// InsightType : Describes insight type.
type InsightType string

const (
	TypeActivity = InsightType("activity")
	TypeSearch   = InsightType("search")
)

// UnmarshalJSON : Custom unmarshal for InsightType.
func (o *InsightType) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))

	switch s {
	case `"activity"`:
		*o = TypeActivity
	case `"search"`:
		*o = TypeSearch
	default:
		return fmt.Errorf("non valid scope option - available options \"activity\"|\"search\"")
	}

	return nil
}
