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
	_ metrics.Metrics = new(Influx)
)

const (
	CPU_MEASUREMENTS = "cpu"
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

func (flux *Influx) CPULastN(n uint, series metrics.SeriesConfig) (vector.Vec, error) {
	var resp, err = flux.Query(
		influx.NewQuery(
			fmt.Sprintf("SELECT value FROM %s ORDER BY time DESC LIMIT %d", CPU_MEASUREMENTS, n),
			flux.dbName,
			"",
		),
	)
	if err != nil {
		return nil, err
	}
	row := resp.Results[0].Series[0]
	ret := vector.MakeVec(len(row.Values), func(index int) float64 {
		f, err := row.Values[index][1].(json.Number).Float64()
		if err != nil {
			panic(err)
		}
		return f
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (flux *Influx) CPUFromTo(from, to time.Time, series metrics.SeriesConfig) (vector.Vec, error) {
	var resp, err = flux.Query(
		influx.NewQuery(
			fmt.Sprintf("SELECT value FROM %s ORDER BY time DESC WHERE time >= %s and time <= %s",
				CPU_MEASUREMENTS, from.Format(time.RFC3339), to.Format(time.RFC3339)),
			flux.dbName,
			"",
		),
	)
	if err != nil {
		return nil, err
	}
	row := resp.Results[0].Series[0]
	ret := vector.MakeVec(len(row.Values), func(index int) float64 {
		f, err := row.Values[index][1].(json.Number).Float64()
		if err != nil {
			panic(err)
		}
		return f
	})
	return ret, nil
}
