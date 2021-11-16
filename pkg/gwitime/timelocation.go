package gwitime

import (
	"database/sql"
	"time"
)

const (
	GWILayout         = "2006-01-02T15:04:05"
	GWILayoutDateOnly = "2006-01-02"

	Athens = "Europe/Athens"
)

// DateTime : Custom gwi time.
type DateTime struct {
	time.Time
}

// Scan : Implements the Scanner interface.
func (o *DateTime) Scan(value interface{}) error {
	nt := sql.NullTime{}
	err := nt.Scan(value)
	if err != nil {
		return err
	}

	if nt.Valid {
		o.Time = nt.Time
	}

	return nil
}

// CreateWithLocation : Returns time based on location.
func CreateWithLocation(t time.Time) DateTime {
	if t.Location().String() == Location().String() {
		return DateTime{Time: t}
	}
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}

// Date : Returns the time.Date result of default local location in DateTime struct.
func Date(year int, month time.Month, day, hour, min, sec, nsec int) DateTime {
	return Create(time.Date(year, month, day, hour, min, sec, nsec, Location()))
}

// Create : Creates from time.
func Create(t time.Time) DateTime {
	return DateTime{Time: t}
}

// Location : Returns default local (Athens) location.
func Location() *time.Location {
	l, _ := time.LoadLocation(Athens)
	return l
}
