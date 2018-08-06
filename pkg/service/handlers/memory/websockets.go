package memory

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

const websocketsStep = 2 * time.Minute

var upgrader = websocket.Upgrader{} // use default options

func HistoryWS(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("START GET metrics MEMORY")
		defer logrus.Debugf("END GET metrics MEMORY")

		fromToStep := handlers.FromToStep{}
		fromToStep.From = time.Now().Add(-1 * time.Hour)
		fromToStep.To = time.Now()
		fromToStep.Step = websocketsStep

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			logrus.Debugf("%+v %d points", fromToStep, fromToStep.To.Sub(fromToStep.From)/fromToStep.Step)
			memoryHistory, err := metrics.MemoryHistory(fromToStep.From, fromToStep.To, fromToStep.Step)
			if err != nil {
				gonic.Gonic(meterrs.ErrUnableToGetMemoryHistory().AddDetailsErr(err), ctx)
				return
			}
			logrus.Debugf("writing response")

			err = c.WriteJSON(memoryHistory)
			if err != nil {
				log.Println("write:", err)
				break
			}
			time.Sleep(websocketsStep)

			fromToStep.From = time.Now()
			fromToStep.To = time.Now()
		}
	}
}

func NodesHistoryWS(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Debugf("START GET nodes metrics MEMORY")
		defer logrus.Debugf("END GET nodes metrics MEMORY")

		fromToStep := handlers.FromToStep{}
		fromToStep.From = time.Now().Add(-1 * time.Hour)
		fromToStep.To = time.Now()
		fromToStep.Step = websocketsStep

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			logrus.Debugf("%+v %d points", fromToStep, fromToStep.To.Sub(fromToStep.From)/fromToStep.Step)
			memoryHistory, err := metrics.MemoryNodesHistory(fromToStep.From, fromToStep.To, fromToStep.Step)
			if err != nil {
				gonic.Gonic(meterrs.ErrUnableToGetMemoryHistory().AddDetailsErr(err), ctx)
				return
			}
			logrus.Debugf("writing response")

			err = c.WriteJSON(memoryHistory)
			if err != nil {
				log.Println("write:", err)
				break
			}
			time.Sleep(websocketsStep)

			fromToStep.From = time.Now()
			fromToStep.To = time.Now()
		}
	}
}
