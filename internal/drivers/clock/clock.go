package clock

import (
	"time"
)

// This is an interface/wrapper around time.Now, used to generate the creation
// times in the usecase. We need this to be able to mock and control the
// generated times during testing.

type Clock interface {
	Now() time.Time
}

type TimeClock struct{}

func NewTimeClock() TimeClock {
	return TimeClock{}
}

func (c TimeClock) Now() time.Time {
	return time.Now()
}
