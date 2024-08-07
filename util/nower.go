package util

import "time"

type Nower interface {
	Now() time.Time
}

type TimeNower struct{}

func NewTimeNower() TimeNower {
	return TimeNower{}
}

func (TimeNower) Now() time.Time {
	return time.Now()
}

type TestNower struct {
	time time.Time
}

func NewTestNower(t time.Time) TestNower {
	return TestNower{
		time: t,
	}
}

func (t TestNower) Now() time.Time {
	return t.time
}
