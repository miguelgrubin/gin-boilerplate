package shared

import (
	"time"
)

type EntityID string

func (e EntityID) AsString() string {
	return string(e)
}

type DateTime time.Time

func (d DateTime) AsTime() time.Time {
	return time.Time(d)
}
