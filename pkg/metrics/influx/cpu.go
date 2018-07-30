package influx

import (
	"encoding/json"
	"log"

	"github.com/containerum/nodeMetrics/pkg/vector"
	"github.com/pkg/errors"
)

var (
	ErrEmptyResult   = errors.New("empty result from InfluxDB")
	ErrNoSeriesFound = errors.New("no series found")
)

func (flux *Influx) CPUCurrent() (uint64, error) {
	var result, err = flux.Query("SELECT value FROM cpu_usage_system LIMIT 1000")
	if err != nil {
		return 0, err
	}
	if len(result) < 1 {
		return 0, ErrEmptyResult
	}
	if len(result[0].Series) < 1 {
		return 0, ErrNoSeriesFound
	}
	var values = result[0].Series[0].Values
	var average = vector.MakeVec(len(values), func(index int) float64 {
		var point = values[index]
		if len(point) < 2 {
			log.Panicf("invalid data point in InfluxDB response: expected >= 2 columns, got %q", point)
		}
		switch point := point[1].(type) {
		case int:
			return float64(point)
		case float64:
			return point
		case json.Number:
			var x, err = point.Float64()
			if err != nil {
				return 0
			}
			return x
		default:
			return 0
		}
	}).DivideScalar(10000000000).Average()
	return uint64(average), nil
}
