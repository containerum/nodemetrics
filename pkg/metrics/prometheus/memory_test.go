package prometheus

import (
	"testing"
	"time"
)

func TestAPI_MemoryHistory(test *testing.T) {
	var client = New(Config{
		Addr:    "http://192.168.88.210:9090",
		Timeout: 10 * time.Second,
	})
	var now = time.Now()
	var data, err = client.MemoryHistory(now.Add(-25*time.Minute), now, 10*time.Second)
	if err != nil {
		test.Fatal(err)
	}
	test.Log(data)
}

func TestAPI_MemoryCurrent(test *testing.T) {
	var client = New(Config{
		Addr:    "http://192.168.88.210:9090",
		Timeout: 10 * time.Second,
	})
	current, err := client.MemoryCurrent()
	if err != nil {
		test.Fatal(err)
	}
	test.Log(current)
}
