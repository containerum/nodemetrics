package metrics

import "testing"

func TestInflux_CPUCurrent(test *testing.T) {
	var db, err = NewInflux(InfluxConfig{
		Database: "kubernetes",
		Addr:     "http://localhost:8888",
	})
	if err != nil {
		test.Fatal(err)
	}
	test.Log(db.MemoryCurrent())
}
