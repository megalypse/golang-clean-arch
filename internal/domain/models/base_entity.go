package models

import "time"

// TODO: Improve filtering when it comes to dates.
type BaseEntity struct {
	Id        *int64     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
