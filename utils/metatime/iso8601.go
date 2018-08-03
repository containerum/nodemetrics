package metatime

import "time"

const (
	ISO8601 = "2006-01-02T15:04:05Z07:00"
)

func NowUTC() time.Time {
	return time.Now().UTC()
}
