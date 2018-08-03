package prometheus

import (
	"testing"
	"time"
)

func TestAPI_StorageCurrent(test *testing.T) {
	var client = New(Config{
		Addr:    "http://192.168.88.210:9090",
		Timeout: 10 * time.Second,
	})
	current, err := client.StorageCurrent()
	if err != nil {
		test.Fatal(err)
	}
	test.Log(current)
}
