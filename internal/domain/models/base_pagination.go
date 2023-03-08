package models

type Paginated[T any] struct {
	Content    []T   `json:"content"`
	Total      int64 `json:"total"`
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	TotalPages int64 `json:"total_pages"`
}
