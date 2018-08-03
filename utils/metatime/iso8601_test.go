package metatime

import (
	"testing"
	"time"
)

func TestISO8601(test *testing.T) {
	var timestamp = "2018-07-31T08:22:56.014Z"
	var t, err = time.Parse(ISO8601, timestamp)
	if err != nil {
		test.Fatal(err)
	}
	test.Log(t)
}
