package influx

import (
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
