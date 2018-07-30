package service

import (
	"time"

	"github.com/containerum/nodeMetrics/pkg/metrics/cadvisor"
	"github.com/containerum/nodeMetrics/pkg/metrics/influx"
	"github.com/containerum/nodeMetrics/pkg/service/handlers/cpu"
	"github.com/containerum/nodeMetrics/pkg/service/handlers/memory"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RubyDate,
	})
	logrus.SetLevel(logrus.DebugLevel)
	var influxStore, err = influx.NewInflux(influx.Config{
		Database: config.DB,
		Addr:     config.InfluxAddr,
	})
	if err != nil {
		return nil, err
	}
	var cadVisor = cadvisor.Mew(config.CadvisorAddr, 60*time.Second)

	var metricsSource = struct {
		*influx.Influx
		*cadvisor.Client
	}{
		Influx: influxStore,
		Client: cadVisor,
	}

	var server = gin.New()
	server.Use(gin.ErrorLogger())
	server.Use(gin.Recovery())

	var CPUmetrics = server.Group("/cpu")
	{
		CPUmetrics.GET("/current", cpu.Current(metricsSource))
	}
	var memoryMetrics = server.Group("/memory")
	{
		memoryMetrics.GET("/current", memory.Current(metricsSource))
	}
	return &Service{
		Engine: server,
		config: config,
	}, nil
}
