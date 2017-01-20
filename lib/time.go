package lib

import "time"

func From1970(microsecsFrom1970 int64) time.Time {
	if microsecsFrom1970 <= 0 {
		return time.Time{}
	}

	secs := microsecsFrom1970 / (1000 * 1000)
	nanosecs := (microsecsFrom1970 - secs * 1000 * 1000) * 1000

	return time.Unix(secs, nanosecs)
}

func MicrosecondsFrom1970(t time.Time) int64 {
	return t.Unix() * 1000 * 1000 + int64(t.Nanosecond()) / 1000
}