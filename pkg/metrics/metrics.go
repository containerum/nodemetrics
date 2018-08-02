package metrics

import (
	"time"

	"github.com/containerum/nodeMetrics/pkg/dataframe"
	"github.com/containerum/nodeMetrics/pkg/models"
)

type Metrics interface {
	CPU
	Memory
	//Storage
}

type CPU interface {
	CPUCurrent() (float64, error)
	CPUHistory(from, to time.Time, step time.Duration) (dataframe.Dataframe, error)
}

type Memory interface {
	MemoryCurrent() (float64, error)
	MemoryHistory(from, to time.Time, step time.Duration) (dataframe.Dataframe, error)
}

type Storage interface {
	StorageCurrent() (models.StorageCurrent, error)
}

func DefaultHistory() (from, to time.Time, step time.Duration) {
	var now = time.Now()
	return now.Add(-12 * time.Hour), now, 15 * time.Minute
}

type Range struct {
	From time.Time
	To   time.Time
	Step time.Duration
}
