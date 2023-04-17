package time

import (
	"fmt"
	"strconv"
)

func ParseUptime(s string) string {
	seconds, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "Unknown"
	}

	minutes := seconds / 60 % 60
	hours := seconds / 60 / 60 % 24
	days := seconds / 60 / 60 / 24

	return fmt.Sprintf("%d day(s), %d hour(s), %d minute(s)", days, hours, minutes)
}
