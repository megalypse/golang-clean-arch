package models

type Person struct {
	Fullname string `json:"fullname"`
	Age      int64  `json:"age"`
	Email    string `json:"email"`
	BaseEntity
}
