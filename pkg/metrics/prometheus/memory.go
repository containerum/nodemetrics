package prometheus

import (
	"fmt"
	"math"
	"time"

	"github.com/containerum/nodeMetrics/pkg/dataframe"
	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/utils/metatime"
	prometheusModel "github.com/prometheus/common/model"
)

var (
	_ metrics.Memory = new(API)
)

func (api *API) MemoryCurrent() (float64, error) {
	var result, err = api.Query(`(node_memory_MemTotal - node_memory_MemFree - node_memory_Buffers- node_memory_Cached)/node_memory_MemTotal`)
	if err != nil {
		return 0, err
	}
	switch data := result.(type) {
	case prometheusModel.Vector:
		for _, sample := range data {
			return math.Round(100 * float64(sample.Value)), nil
		}
	default:
		return 0, fmt.Errorf("[prometheus.API.MemoryCurrent] unexpected value type %q", data.Type())
	}
	return 0, nil
}

func (api *API) MemoryHistory(from, to time.Time, step time.Duration) (dataframe.Dataframe, error) {
	var result, err = api.QueryRange(metrics.Range{
		From: from,
		To:   to,
		Step: step,
	}, `(node_memory_MemTotal - node_memory_MemFree - node_memory_Buffers- node_memory_Cached)/node_memory_MemTotal`)
	if err != nil {
		return dataframe.Dataframe{}, err
	}
	switch data := result.(type) {
	case prometheusModel.Matrix:
		if data.Len() < 1 {
			return dataframe.Dataframe{}, nil
		}
		var points = data[0].Values
		var history = dataframe.MakeDataframe("%", len(points), func(index int) (string, float64) {
			var point = points[index]
			var timestamp = point.Timestamp.Time().Format(metatime.ISO8601)
			var value = math.Round(100 * float64(point.Value))
			return timestamp, value
		})
		return history, nil
	default:
		return dataframe.Dataframe{}, fmt.Errorf("[prometheus.API.MemoryHistory] unexpected value type %q", data.Type())
	}
	return dataframe.Dataframe{}, nil
}
