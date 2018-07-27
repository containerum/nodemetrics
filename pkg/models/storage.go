package models

type StorageCurrent struct {
	Units   string `json:"units"`
	Storage uint64 `json:"storage"`
	Limit   uint64 `json:"limit"`
}
