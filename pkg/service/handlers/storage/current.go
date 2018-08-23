package storage

import (
	"net/http"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/nodeMetrics/pkg/meterrs"
	"github.com/containerum/nodeMetrics/pkg/metrics"
	"github.com/containerum/nodeMetrics/pkg/models"
	"github.com/gin-gonic/gin"
)

var (
	_ gin.HandlerFunc = Current(nil)
)

func Current(metrics metrics.Storage) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var storageUsage, err = metrics.StorageCurrent()
		if err != nil {
			gonic.Gonic(meterrs.ErrUnableToGetMemoryCurrent().AddDetailsErr(err), ctx)
			return
		}
		ctx.JSON(http.StatusOK, models.StorageCurrent{
			Units:   "%",
			Storage: storageUsage,
		})
	}
}
