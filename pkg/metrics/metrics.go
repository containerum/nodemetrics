package metrics

import "github.com/containerum/nodeMetrics/pkg/models"

type Metrics interface {
	CPU
	Memory
	Storage
}

type CPU interface {
	CPUCurrent() (uint64, error)
}

type Memory interface {
	MemoryCurrent() (uint64, error)
}

type Storage interface {
	StorageCurrent() (models.StorageCurrent, error)
}
