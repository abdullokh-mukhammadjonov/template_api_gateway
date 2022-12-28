package util

import (
	"errors"
	"strings"
	"time"
)

func IsNullDate(dateString string) (bool, error) {
	if len(dateString) < 10 {
		return false, errors.New("date string must be at least 10 characters. e.g.: 2022-05-23T01:03:16.000+00:00")
	}
	return dateString[:10] == "0001-01-01", nil
}

func TimeParseWithCheck(layout string, timeString string) (time.Time, error) {
	parsed, err := time.Parse(layout, timeString)
	if err != nil {
		if strings.Contains(err.Error(), "month out of range") {
			dateParts := strings.Split(timeString, "-")
			if len(dateParts) != 3 {
				dateParts = strings.Split(timeString, "/")
			}

			// switching date parts
			// example 2022-17-05   to   2022-05-17
			second := dateParts[1]
			dateParts[1] = dateParts[2]
			dateParts[2] = second

			reformed := strings.Join(dateParts, "-")
			parsed, err = time.Parse(layout, reformed)
			if err != nil {
				return time.Time{}, err
			}
		} else {
			return time.Time{}, err
		}
	}
	return parsed, nil
}
