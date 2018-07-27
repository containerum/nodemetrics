package metrics

type Metrics interface {
	CPU
	Memory
}

type CPU interface {
	CPUCurrent() (uint64, error)
}

type Memory interface {
	MemoryCurrent() (uint64, error)
}
