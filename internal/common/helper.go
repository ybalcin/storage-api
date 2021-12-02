// Package common provides utilities
package common

import (
	"errors"
	"time"
)

// TimeToDateOnlyFormat returns time as date only format
func TimeToDateOnlyFormat(t time.Time) (time.Time, error) {
	return time.Parse("2006-01-02", t.String())
}

// StringToTime returns string as time
func StringToTime(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, errors.New("time is empty")
	}

	return time.Parse("2006-01-02", t)
}
