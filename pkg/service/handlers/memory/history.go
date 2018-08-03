package memory

import (
	"net/http"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/nodeMetrics/pkg/meterrs"
	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func History(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("START GET metrics history")
		defer logrus.Debugf("END GET metrics history")

		fromToStep, parsingErr := handlers.ParseFromToStep(ctx)
		if parsingErr != nil {
			gonic.Gonic(parsingErr, ctx)
			return
		}
		logrus.Debugf("%+v %d points", fromToStep, fromToStep.To.Sub(fromToStep.From)/fromToStep.Step)
		memoryHistory, err := metrics.MemoryHistory(fromToStep.From, fromToStep.To, fromToStep.Step)
		if err != nil {
			gonic.Gonic(meterrs.ErrUnableToGetMemoryHistory().AddDetailsErr(err), ctx)
			return
		}
		logrus.Debugf("writing response")
		ctx.JSON(http.StatusOK, memoryHistory)
	}
}

func NodeHistory(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("START GET nodes metrics history")
		defer logrus.Debugf("END GET nodes metrics history")

		fromToStep, parsingErr := handlers.ParseFromToStep(ctx)
		if parsingErr != nil {
			gonic.Gonic(parsingErr, ctx)
			return
		}
		logrus.Debugf("%+v %d points", fromToStep, fromToStep.To.Sub(fromToStep.From)/fromToStep.Step)
		memoryHistory, err := metrics.MemoryNodesHistory(fromToStep.From, fromToStep.To, fromToStep.Step)
		if err != nil {
			gonic.Gonic(meterrs.ErrUnableToGetMemoryHistory().AddDetailsErr(err), ctx)
			return
		}
		logrus.Debugf("writing response")
		ctx.JSON(http.StatusOK, memoryHistory)
	}
}
