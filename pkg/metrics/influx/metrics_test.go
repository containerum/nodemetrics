package influx

import (
	"fmt"
	"testing"
	"time"

	"net/url"

	"github.com/containerum/nodeMetrics/utils/metatime"
)

func TestInflux_CPUCurrent(test *testing.T) {
	var db, err = NewInflux(Config{
		Database: "kubernetes",
		Addr:     "http://192.168.88.210:8086",
		NumCPU:   4,
	})
	if err != nil {
		test.Fatal(err)
	}
	var now = time.Now()
	fmt.Println(url.Values{
		"from": []string{now.Add(-72 * time.Hour).Format(metatime.ISO8601)},
		"to":   []string{now.Add(-48 * time.Hour).Format(metatime.ISO8601)},
	}.Encode())
	test.Log(db.MemoryHistory(now.Add(-72*time.Hour), now.Add(-48*time.Hour), 15*time.Minute))
	// test.Log(db.MemoryCurrent())
}
