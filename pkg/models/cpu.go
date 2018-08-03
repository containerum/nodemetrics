package models

type CPUCurrent struct {
	Units string  `json:"units"`
	CPU   float64 `json:"cpu"`
}

type CPUHistory struct {
	Units  string    `json:"units"`
	Memory []float64 `json:"memory"`
}
