package memory

import (
	"net/http"

	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/models"
	"github.com/gin-gonic/gin"
)

var (
	_ gin.HandlerFunc = Current(nil)
)

func Current(metrics metrics.Metrics) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var cpuMetrics, err = metrics.MemoryCurrent()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, models.CPUCurrent{
			Units: "Mb",
			CPU:   cpuMetrics,
		})
	}
}