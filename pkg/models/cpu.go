package models

type CPUCurrent struct {
	Units string `json:"units"`
	CPU   uint64 `json:"cpu"`
}

type MemoryCurrent struct {
	Units  string `json:"units"`
	Memory uint64 `json:"memory"`
}
