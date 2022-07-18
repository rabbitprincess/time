package time

import (
	"fmt"
	"time"
)

func NewTimeDiff() *TimeDiff {
	c := &TimeDiff{}
	return c
}

type TimeDiff struct {
	timeStart time.Time
	duEnd     time.Duration
}

func (t *TimeDiff) Start() {
	t.timeStart = UTC_now()
}

func (t *TimeDiff) End() {
	t.duEnd = UTC_now().Sub(t.timeStart)
}

func (t *TimeDiff) PrintDiff() {
	fmt.Printf("%s\n", t.duEnd.String())
}
func (t *TimeDiff) GetDiff_str() string {
	return t.duEnd.String()
}

func (t *TimeDiff) GetDiff_nano() int64 {
	return t.duEnd.Nanoseconds()
}

func (t *TimeDiff) GetDiff_micro() int64 {
	return t.duEnd.Microseconds()
}

func (t *TimeDiff) GetDiff_milli() int64 {
	return t.duEnd.Milliseconds()
}

func (t *TimeDiff) GetDiff_sec() int64 {
	return int64(t.duEnd.Seconds())
}
