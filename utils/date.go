package utils

import (
	e "application-design-test-master/errors"
	"time"
)

const (
	hourInDay = 24
)

// DaysBetween return days between two arguments.
func DaysBetween(from time.Time, to time.Time) ([]time.Time, error) {
	if from.After(to) {
		return nil, e.ErrInvalidDateRange
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days, nil
}

func toDay(timestamp time.Time) time.Time {
	return timestamp.Truncate(hourInDay * time.Hour)
}

// Date convert arguments to date in time.Time.
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
