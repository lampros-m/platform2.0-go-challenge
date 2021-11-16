package gwitime

import (
	"encoding/json"
	"strings"
	"time"
)

// UnmarshalJSON function.
func (dt *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(strings.TrimSpace(string(b)), "\"")

	t, err := time.Parse(GWILayout, s)
	if err != nil {
		t, err = time.Parse(GWILayoutDateOnly, s)
		if err != nil {
			return err
		}
	}

	*dt = CreateWithLocation(t)
	return nil
}

// MarshalJSON function.
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.IsZero() {
		return []byte(`""`), nil
	}
	str := dt.Time.Format(GWILayout)

	return json.Marshal(str)
}
