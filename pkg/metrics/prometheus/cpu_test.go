package prometheus

import (
	"testing"
	"time"
)

func TestAPI_CPUHistory(test *testing.T) {
	var client = New(Config{
		Addr:    "http://192.168.88.210:9090",
		Timeout: 10 * time.Second,
	})
	var now = time.Now()
	var data, err = client.CPUHistory(now.Add(-15*time.Minute), now, 10*time.Second)
	if err != nil {
		test.Fatal(err)
	}
	test.Log(data)
}

func TestAPI_CPUCurrent(test *testing.T) {
	var client = New(Config{
		Addr:    "http://192.168.88.210:9090",
		Timeout: 10 * time.Second,
	})
	current, err := client.CPUCurrent()
	if err != nil {
		test.Fatal(err)
	}
	test.Log(current)
}
