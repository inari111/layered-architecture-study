package domain

import "time"

type CurrentTimeFunc func() time.Time

func NewCurrentTimeFunc() CurrentTimeFunc {
	return time.Now
}
