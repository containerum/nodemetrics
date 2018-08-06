package cpu

import (
	"log"
	"time"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/nodeMetrics/pkg/meterrs"
	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{} // use default options

func HistoryWS(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("START GET metrics CPU")
		defer logrus.Debugf("END GET metrics CPU")

		fromToStep := handlers.FromToStep{}
		fromToStep.From = time.Now().Add(-1 * time.Hour)
		fromToStep.To = time.Now()
		fromToStep.Step = 1 * time.Minute

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			logrus.Debugf("%+v %d points", fromToStep, fromToStep.To.Sub(fromToStep.From)/fromToStep.Step)
			cpuHistory, err := metrics.CPUHistory(fromToStep.From, fromToStep.To, fromToStep.Step)
			if err != nil {
				gonic.Gonic(meterrs.ErrUnableToGetCPUHistory().AddDetailsErr(err), ctx)
				return
			}
			logrus.Debugf("writing response")

			err = c.WriteJSON(cpuHistory)
			if err != nil {
				log.Println("write:", err)
				break
			}
			time.Sleep(1 * time.Minute)

			fromToStep.From = time.Now()
			fromToStep.To = time.Now()
			fromToStep.Step = 1 * time.Minute
		}
	}
}

func NodesHistoryWS(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("START GET nodes metrics CPU")
		defer logrus.Debugf("END GET nodes metrics CPU")

		fromToStep := handlers.FromToStep{}
		fromToStep.From = time.Now().Add(-1 * time.Hour)
		fromToStep.To = time.Now()
		fromToStep.Step = 1 * time.Minute

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			logrus.Debugf("%+v %d points", fromToStep, fromToStep.To.Sub(fromToStep.From)/fromToStep.Step)
			cpuHistory, err := metrics.CPUNodesHistory(fromToStep.From, fromToStep.To, fromToStep.Step)
			if err != nil {
				gonic.Gonic(meterrs.ErrUnableToGetCPUHistory().AddDetailsErr(err), ctx)
				return
			}
			logrus.Debugf("writing response")

			err = c.WriteJSON(cpuHistory)
			if err != nil {
				log.Println("write:", err)
				break
			}
			time.Sleep(1 * time.Minute)

			fromToStep.From = time.Now()
			fromToStep.To = time.Now()
			fromToStep.Step = 1 * time.Minute
		}
	}
}
