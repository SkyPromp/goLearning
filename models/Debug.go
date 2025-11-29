package models

type Debug struct {
	Duration int64 `json:"duration_ns"`
	Data any `json:"data"`
}
