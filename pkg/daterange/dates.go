package daterange

import (
	"time"

	"gwi/platform2.0-go-challenge/pkg/gwitime"
)

// DateRange : Date range struct for requests.
type DateRange struct {
	DateFrom gwitime.DateTime `json:"date_from"`
	DateTo   gwitime.DateTime `json:"date_to"`
}

// GetDefault : Returns default date range.
func GetDefault() DateRange {
	dateRange := DateRange{
		DateFrom: gwitime.Create(time.Now().AddDate(0, 0, -15)),
		DateTo:   gwitime.Create(time.Now()),
	}

	return dateRange
}
