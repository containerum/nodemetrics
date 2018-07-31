package handlers

import (
	"net/http"
	"time"

	"github.com/containerum/nodeMetrics/pkg/models"
	"github.com/containerum/nodeMetrics/utils/metatime"
	"github.com/gin-gonic/gin"
)

type FromToStep struct {
	From, To time.Time
	Step     time.Duration
}

type Error struct {
	Status int
	models.Error
}

func ErrorF(status int, msg string, args ...interface{}) *Error {
	return &Error{
		Status: status,
		Error:  *models.ErrorF(msg, args...),
	}
}

func (err *Error) ToGin(ctx *gin.Context) {
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status, err.Error)
	}
}

func ParseFromToStep(ctx *gin.Context) (FromToStep, *Error) {
	var to = time.Now()
	var from = to.Add(-12 * time.Hour)
	var step = 15 * time.Minute

	if fromStr := ctx.Query("from"); fromStr != "" {
		var err error
		from, err = time.Parse(metatime.ISO8601, fromStr)
		if err != nil {
			return FromToStep{}, ErrorF(http.StatusBadRequest, "invalid 'from' parameter format: ISO8601 is expected, got %q", fromStr)
		}
	}

	if toStr := ctx.Query("to"); toStr != "" {
		var err error
		to, err = time.Parse(metatime.ISO8601, toStr)
		if err != nil {
			return FromToStep{}, ErrorF(http.StatusBadRequest, "invalid 'to' parameter format: ISO8601 is expected, got %q", toStr)
		}
	}

	if stepStr := ctx.Query("step"); stepStr != "" {
		var err error
		step, err = time.ParseDuration(stepStr)
		if err != nil {
			return FromToStep{}, ErrorF(http.StatusBadRequest, "invalid 'step' parameter format: A duration string is a possibly signed sequence of "+
				"decimal numbers, each with optional fraction and a unit suffix, "+
				`such as "300ms", "-1.5h" or "2h45m". `+
				`Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
		}
	}

	if !from.Before(to) {
		return FromToStep{}, ErrorF(http.StatusBadRequest, "'from' timestamp must be before 'to' timestamp")
	}

	if step < 10*time.Second {
		return FromToStep{}, ErrorF(http.StatusBadRequest, "'step' parameter must be > then 10s")
	}
	if to.Sub(from)/step > 1000000 {
		return FromToStep{}, ErrorF(http.StatusRequestEntityTooLarge, "can't server more then 1000000 data points")
	}

	return FromToStep{
		From: from,
		To:   to,
		Step: step,
	}, nil
}
