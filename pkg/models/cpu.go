package models

type CPUCurrent struct {
	Units string `json:"units"`
	CPU   uint64 `json:"cpu"`
}
