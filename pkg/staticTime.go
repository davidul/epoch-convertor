package pkg

import "time"

type StaticTime interface {
	Now() time.Time
}

type RealTime struct{}
type MockTime struct{}

func (*MockTime) Now() time.Time {
	return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
}

func (*RealTime) Now() time.Time {
	return time.Now()
}
