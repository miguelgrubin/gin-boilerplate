package shared

import (
	"time"
)

type EntityId string

func (e EntityId) AsString() string {
	return string(e)
}

type DateTime time.Time

func (d DateTime) AsTime() time.Time {
	return time.Time(d)
}
