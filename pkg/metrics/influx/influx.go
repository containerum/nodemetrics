package influx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/vector"
	influx "github.com/influxdata/influxdb/client/v2"
)

var (
	_ interface {
		metrics.CPU
		metrics.Memory
	} = new(Influx)
)

type Influx struct {
	dbName string
	influx.Client
}

type Config struct {
	Database string
	Addr     string
	Username string
	Password string
}

func NewInflux(config Config) (*Influx, error) {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     config.Addr,
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		return nil, err
	}
	return &Influx{
		Client: client,
		dbName: config.Database,
	}, nil
}

func (flux *Influx) CPUFromTo(from, to time.Time, series SeriesConfig) (vector.Vec, error) {
	panic("NOT IMPLEMENTED")
	return nil, nil
}

func (flux *Influx) CPUCurrent() (uint64, error) {
	var result, err = flux.Query("SELECT value FROM cpu_usage_system WHERE time > now()-80s LIMIT 1000")
	if err != nil {
		return 0, err
	}
	var points = result[0].Series[0].Values
	var average = vector.MakeVec(len(points), func(index int) float64 {
		switch point := points[index][1].(type) {
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
			fmt.Printf("%T %v\n", point, point)
			return 0
		}
	}).DivideScalar(10000000000).Average()
	return uint64(average), nil
}

func (flux *Influx) MemoryCurrent() (uint64, error) {
	var result, err = flux.Query("SELECT value FROM memory_usage WHERE time > now()-80s LIMIT 1000")
	if err != nil {
		return 0, err
	}
	var points = result[0].Series[0].Values
	var average = vector.MakeVec(len(points), func(index int) float64 {
		switch point := points[index][1].(type) {
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
			fmt.Printf("%T %v\n", point, point)
			return 0
		}
	}).DivideScalar(1000).Average()
	return uint64(average), nil
}

func (flux *Influx) StorageCurrent() (uint64, error) {
	var result, err = flux.Query("SELECT value FROM fs_usage WHERE time > now()-80s LIMIT 1000")
	if err != nil {
		return 0, err
	}
	var points = result[0].Series[0].Values
	var average = vector.MakeVec(len(points), func(index int) float64 {
		switch point := points[index][1].(type) {
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
			fmt.Printf("%T %v\n", point, point)
			return 0
		}
	}).DivideScalar(1).Average()
	return uint64(average), nil
}
