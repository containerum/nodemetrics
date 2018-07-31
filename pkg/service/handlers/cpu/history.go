package cpu

import (
	"net/http"

	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/models"
	"github.com/containerum/nodeMetrics/pkg/service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func History(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("START GET metrics CPU")
		defer logrus.Debugf("END GET metrics CPU")

		fromToStep, parsingErr := handlers.ParseFromToStep(ctx)
		if parsingErr != nil {
			parsingErr.ToGin(ctx)
			return
		}
		logrus.Debugf("%+v", fromToStep)
		memoryHistory, err := metrics.CPUHistory(fromToStep.From, fromToStep.To, fromToStep.Step)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, models.MemoryHistory{
			Units:  "%",
			Memory: memoryHistory,
		})
	}
}
