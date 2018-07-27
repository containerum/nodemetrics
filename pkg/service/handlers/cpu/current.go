package cpu

import (
	"net/http"

	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	_ gin.HandlerFunc = Current(nil)
)

func Current(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("serving current CPU metrics")
		var cpuMetrics, err = metrics.CPUCurrent()
		if err != nil {
			logrus.WithError(err).Errorf("unable to get current CPU metrics")
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, models.CPUCurrent{
			Units: "%",
			CPU:   cpuMetrics,
		})
	}
}
