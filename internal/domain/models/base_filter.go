package models

type BaseFilter struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
