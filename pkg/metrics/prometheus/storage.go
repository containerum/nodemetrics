package prometheus

import (
	"fmt"
	"math"

	"github.com/containerum/nodeMetrics/pkg/metrics"
	prometheusModel "github.com/prometheus/common/model"
)

var (
	_ metrics.Storage = new(API)
)

func (api *API) StorageCurrent() (float64, error) {
	var result, err = api.Query(`100*(sum(node_filesystem_size_bytes{device!="rootfs"}) - sum(node_filesystem_free_bytes{device!="rootfs"})) / sum(node_filesystem_size_bytes{device!="rootfs"})`)
	if err != nil {
		return 0, err
	}
	switch data := result.(type) {
	case prometheusModel.Vector:
		for _, sample := range data {
			return math.Round(float64(sample.Value)), nil
		}
	default:
		return 0, fmt.Errorf("[prometheus.API.StorageCurrent] unexpected value type %q", data.Type())
	}
	return 0, nil
}
