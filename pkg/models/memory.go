package models

type MemoryCurrent struct {
	Units  string `json:"units"`
	Memory uint64 `json:"memory"`
}

type MemoryHistory struct {
	Units  string    `json:"units"`
	Memory []float64 `json:"memory"`
}
