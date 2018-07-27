package main

import (
	"os"

	"github.com/containerum/nodeMetrics/pkg/service"
	"github.com/octago/sflags/gen/gflag"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServingAddr string
	service.Config
}

func main() {
	var config = Config{
		ServingAddr: "localhost:8090",
		Config: service.Config{
			DB:         "kubernetes",
			InfluxAddr: "http://localhost:8888",
		},
	}
	if _, err := gflag.Parse(&config); err != nil {
		panic(err)
	}
	var service, err = service.NewService(config.Config)
	if err != nil {
		logrus.WithError(err).Errorf("unable to start service")
		os.Exit(1)
	}
	if err := service.Run(config.ServingAddr); err != nil {
		logrus.WithError(err).Errorf("error while service execution")
		os.Exit(1)
	}
}
