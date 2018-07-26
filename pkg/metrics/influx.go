package metrics

import (
	"time"

	"github.com/containerum/nodeMetrics/pkg/vector"
	influx "github.com/influxdata/influxdb/client/v2"
)

var (
	_ Metrics = new(Influx)
)

const (
	CPU_COLUMN = "cpu"
)

type Influx struct {
	dbName string
	influx.Client
}

type InfluxConfig struct {
	Database string
	Addr     string
	Username string
	Password string
}

func NewInflux(config InfluxConfig) (*Influx, error) {
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

func (flux *Influx) CPULastN(n uint, series SeriesConfig) (vector.Vec, error) {
	var resp, err = flux.Query(flux.SelectFromWithLimit(CPU_COLUMN, n, "*"))
	if err != nil {
		return nil, err
	}
	_ = resp.Results
	return nil, nil
}

func (flux *Influx) CPUFromTo(from, to time.Time, series SeriesConfig) (vector.Vec, error) {
	panic("NOT IMPLEMENTED")
	return nil, nil
}
