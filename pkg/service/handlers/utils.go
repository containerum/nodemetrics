package handlers

import (
	"time"

	"github.com/containerum/cherry"
	"github.com/containerum/nodeMetrics/pkg/meterrs"
	"github.com/containerum/nodeMetrics/utils/metatime"
	"github.com/gin-gonic/gin"
)

const (
	maxDataponts = 1000000
)

type FromToStep struct {
	From, To time.Time
	Step     time.Duration
}

func ParseFromToStep(ctx *gin.Context) (FromToStep, *cherry.Err) {
	var to = time.Now()
	var from = to.Add(-12 * time.Hour)
	var step = 15 * time.Minute

	if fromStr := ctx.Query("from"); fromStr != "" {
		var err error
		from, err = time.Parse(metatime.ISO8601, fromStr)
		if err != nil {
			return FromToStep{}, meterrs.ErrInvalidQueryParameter().AddDetailF("invalid 'from' parameter format: ISO8601 is expected, got %q", fromStr)
		}
	}

	if toStr := ctx.Query("to"); toStr != "" {
		var err error
		to, err = time.Parse(metatime.ISO8601, toStr)
		if err != nil {
			return FromToStep{}, meterrs.ErrInvalidQueryParameter().AddDetailF("invalid 'to' parameter format: ISO8601 is expected, got %q", toStr)
		}
	}

	if stepStr := ctx.Query("step"); stepStr != "" {
		var err error
		step, err = time.ParseDuration(stepStr)
		if err != nil {
			return FromToStep{}, meterrs.ErrInvalidQueryParameter().AddDetailF("invalid 'step' parameter format: A duration string is a possibly signed sequence of " +
				"decimal numbers, each with optional fraction and a unit suffix, " +
				`such as "300ms", "-1.5h" or "2h45m". ` +
				`Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
		}
	}

	if !from.Before(to) {
		return FromToStep{}, meterrs.ErrInvalidQueryParameter().AddDetailF("'from' timestamp must be before 'to' timestamp")
	}

	if step < 10*time.Second {
		return FromToStep{}, meterrs.ErrInvalidQueryParameter().AddDetailF("'step' parameter must be > then 10s")
	}

	if to.Sub(from)/step > maxDataponts {
		return FromToStep{}, meterrs.ErrTooMuchDataPointsToCalculate().AddDetailF("can't server more then %d data points", maxDataponts)
	}

	return FromToStep{
		From: from,
		To:   to,
		Step: step,
	}, nil
}
