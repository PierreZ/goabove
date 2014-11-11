package goabove

import (
	"encoding/json"
	"time"
)

// Datetimes
// DateTime represents date as returned by Runabove
type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//2014-09-16T06:50:09+02:00 RFC3339
	//t, err := time.Parse("2006-01-02T15:04:05+02:00", s)
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}
