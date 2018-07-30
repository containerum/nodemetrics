package metrics

import (
	"time"

	"github.com/containerum/nodeMetrics/pkg/models"
	"github.com/containerum/nodeMetrics/pkg/vector"
)

type Metrics interface {
	CPU
	Memory
	Storage
}

type CPU interface {
	CPUCurrent() (uint64, error)
	CPUHistory(from, to time.Time, step time.Duration) (vector.Vec, error)
}

type Memory interface {
	MemoryCurrent() (uint64, error)
}

type Storage interface {
	StorageCurrent() (models.StorageCurrent, error)
}

func DefaultHistory() (from, to time.Time, step time.Duration) {
	var now = time.Now()
	return now.Add(-12 * time.Hour), now, 15 * time.Minute
}
