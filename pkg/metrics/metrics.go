package metrics

import (
	"time"

	"github.com/containerum/nodeMetrics/pkg/vector"
)

type Metrics interface {
	CPU
}

type CPU interface {
	CPULastN(n uint, precision SeriesConfig) (vector.Vec, error)
	CPUFromTo(from, to time.Time, precision SeriesConfig) (vector.Vec, error)
}
