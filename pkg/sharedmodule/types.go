package sharedmodule

import (
	"time"
)

type EntityID string

func (e EntityID) String() string {
	return string(e)
}

type DateTime time.Time

func (d DateTime) Time() time.Time {
	return time.Time(d)
}
