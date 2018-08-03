package influx

import (
	"github.com/containerum/nodeMetrics/pkg/metrics"
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
	numCPU uint64
	memory uint64
}

type Config struct {
	Database string
	Addr     string
	Username string
	Password string
	NumCPU   uint64
	Memory   uint64
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
	if config.NumCPU == 0 {
		config.NumCPU = 1
	}
	if config.Memory == 0 {
		config.Memory = 4 * 10 << 10
	}
	return &Influx{
		Client: client,
		dbName: config.Database,
		numCPU: config.NumCPU,
		memory: config.Memory,
	}, nil
}

func (flux *Influx) CPUFactor() float64 {
	return float64(flux.numCPU)
}

func (flux *Influx) MemoryFactor() float64 {
	return float64(flux.memory)
}
