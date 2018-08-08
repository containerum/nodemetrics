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
	_ metrics.CPU = new(API)
)

func (api *API) CPUCurrent() (float64, error) {
	var result, err = api.Query(`100 - avg((irate(node_cpu{mode="idle"}[5m])) * 100)`)
	if err != nil {
		return 0, err
	}
	switch data := result.(type) {
	case prometheusModel.Vector:
		for _, sample := range data {
			return float64(sample.Value), nil
		}
	default:
		return 0, fmt.Errorf("[prometheus.API.CPUCurrent] unexpected value type %q", data.Type())
	}
	return 0, nil
}

func (api *API) CPUHistory(from, to time.Time, step time.Duration) (dataframe.Dataframe, error) {
	var result, err = api.QueryRange(metrics.Range{
		From: from,
		To:   to,
		Step: step,
	}, `100 - avg(irate(node_cpu{mode="idle"}[5m]) * 100)`)
	if err != nil {
		return dataframe.Dataframe{}, err
	}
	switch data := result.(type) {
	case prometheusModel.Matrix:
		if data.Len() < 1 {
			return dataframe.MakeDataframe("%", 0, func(index int) (string, float64) { return "", 0 }), nil
		}
		var points = data[0].Values
		var history = dataframe.MakeDataframe("%", len(points), func(index int) (string, float64) {
			var point = points[index]
			var timestamp = point.Timestamp.Time().Format(metatime.ISO8601)
			var value = math.Round(float64(point.Value))
			return timestamp, value
		})
		return history, nil
	default:
		return dataframe.Dataframe{}, fmt.Errorf("[prometheus.API.CPUHistory] unexpected value type %q", data.Type())
	}
}

func (api *API) CPUNodesHistory(from, to time.Time, step time.Duration) (map[string]dataframe.Dataframe, error) {
	ret := make(map[string]dataframe.Dataframe)
	var result, err = api.QueryRange(metrics.Range{
		From: from,
		To:   to,
		Step: step,
	}, `100 - (avg by (instance) (irate(node_cpu{mode="idle"}[5m])) * 100)`)
	if err != nil {
		return ret, err
	}
	switch data := result.(type) {
	case prometheusModel.Matrix:
		for k := range data {
			var points = data[k].Values
			var history = dataframe.MakeDataframe("%", len(points), func(index int) (string, float64) {
				var point = points[index]
				var timestamp = point.Timestamp.Time().Format(metatime.ISO8601)
				var value = math.Round(float64(point.Value))
				return timestamp, value
			})
			ret[string(data[k].Metric["instance"])] = history
		}
		return ret, nil
	default:
		return ret, fmt.Errorf("[prometheus.API.CPUHistory] unexpected value type %q", data.Type())
	}
}
