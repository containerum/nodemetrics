package influx

import (
	"testing"
)

func TestInflux_CPUCurrent(test *testing.T) {
	var db, err = NewInflux(Config{
		Database: "kubernetes",
		Addr:     "http://192.168.88.210:8086",
	})
	if err != nil {
		test.Fatal(err)
	}
	// var now = time.Now()
	// test.Log(db.CPUHistory(now.Add(-48*time.Hour), now.Add(-24*time.Hour), 15*time.Minute))
	test.Log(db.MemoryCurrent())
}
