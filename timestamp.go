package timestamp

import (
	"errors"
	"time"
)

// TimeStamp should be updated automatically every time
type TimeStamp struct {
	Created time.Time
	Updated time.Time
}

// IsZero checks zero time
func (ts *TimeStamp) IsZero() bool {
	return ts.Created.IsZero() && ts.Updated.IsZero()
}

// Update implements Updater
func (ts *TimeStamp) Update() error {
	now := time.Now()
	if ts.Created.IsZero() {
		return errors.New("timestamp: created is zero")
	}

	ts.Updated = now

	return nil
}
