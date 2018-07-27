package service

import (
	"time"

	"github.com/containerum/nodeMetrics/pkg/metrics/cadvisor"
	"github.com/containerum/nodeMetrics/pkg/metrics/influx"
	"github.com/containerum/nodeMetrics/pkg/service/handlers/cpu"
	"github.com/containerum/nodeMetrics/pkg/service/handlers/memory"
	"github.com/gin-gonic/gin"
)

type Config struct {
	InfluxAddr   string
	CadvisorAddr string
	Username     string
	Password     string
	DB           string
}

type Service struct {
	*gin.Engine
	config Config
}

func NewService(config Config) (*Service, error) {

	var influxStore, err = influx.NewInflux(influx.Config{
		Database: config.DB,
		Addr:     config.InfluxAddr,
	})
	if err != nil {
		return nil, err
	}
	var cadVisor = cadvisor.Mew(config.CadvisorAddr, 60*time.Second)

	var store = struct {
		*influx.Influx
		*cadvisor.Client
	}{
		Influx: influxStore,
		Client: cadVisor,
	}

	var server = gin.New()
	var CPUmetrics = server.Group("/cpu")
	{
		CPUmetrics.GET("/current", cpu.Current(store))
	}
	var memoryMetrics = server.Group("/memory")
	{
		memoryMetrics.GET("/current", memory.Current(store))
	}
	return &Service{
		Engine: server,
		config: config,
	}, nil
}
