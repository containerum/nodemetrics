package service

import (
	"github.com/containerum/nodeMetrics/pkg/metrics/influx"
	"github.com/containerum/nodeMetrics/pkg/service/handlers/cpu"
	"github.com/containerum/nodeMetrics/pkg/service/handlers/memory"
	"github.com/gin-gonic/gin"
)

type Config struct {
	InfluxAddr string
	Username   string
	Password   string
	DB         string
}

type Service struct {
	*gin.Engine
	config Config
}

func NewService(config Config) (*Service, error) {

	var metrics, err = influx.NewInflux(influx.Config{
		Database: config.DB,
		Addr:     config.InfluxAddr,
	})
	if err != nil {
		return nil, err
	}

	var server = gin.New()
	var CPUmetrics = server.Group("/cpu")
	{
		CPUmetrics.GET("/current", cpu.Current(metrics))
	}
	var memoryMetrics = server.Group("/memory")
	{
		memoryMetrics.GET("/current", memory.Current(metrics))
	}
	return &Service{
		Engine: server,
		config: config,
	}, nil
}
