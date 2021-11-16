package charts

// VisitsPerPeriod : Describes visits per period.
type VisitsPerPeriod struct {
	Period string `json:"period"`
	Visits uint32 `json:"visits"`
}

// VisitsChart : Describes visits chart.
type VisitsChart struct {
	AssetID         uint32 `json:"-"`
	VisitsPerPeriod []VisitsPerPeriod
}

// AudienceReachPerPeriod : Describes audience reached per period.
type AudienceReachPerPeriod struct {
	Period  string `json:"period"`
	Reached uint32 `json:"reached"`
}

// AudienceReachChart : Audience reach chart.
type AudienceReachChart struct {
	AssetID                uint32 `json:"-"`
	AudienceReachPerPeriod []AudienceReachPerPeriod
}
