package cadvisor

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/models"
)

var (
	_ metrics.Storage = &Client{}
)

type Client struct {
	addr string
	*http.Client
}

func Mew(addr string, timeout time.Duration) *Client {
	return &Client{
		addr: addr,
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (client *Client) Report() (models.Report, error) {
	var resp, err = http.Get(client.addr + "/api/v1.0/containers/")
	if err != nil {
		return models.Report{}, err
	}
	defer resp.Body.Close()
	var report models.Report
	return report, json.NewDecoder(resp.Body).Decode(&report)
}

func (client *Client) StorageCurrent() (models.StorageCurrent, error) {
	var report, err = client.Report()
	if err != nil {
		return models.StorageCurrent{}, err
	}
	var storage models.StorageCurrent
	for _, stat := range report.Stats {
		for _, fs := range stat.Filesystem {
			storage.Storage += uint64(fs.Usage)
			storage.Limit += uint64(fs.Capacity)
		}
	}
	return storage, nil
}
