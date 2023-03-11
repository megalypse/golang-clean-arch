package models

type BaseFilter[T any] struct {
	Filters T   `json:"filters"`
	Page    int `json:"page"`
	Limit   int `json:"limit"`
}
