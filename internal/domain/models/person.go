package models

type Person struct {
	Fullname string `json:"fullname"`
	Age      int8   `json:"age"`
	Email    string `json:"email"`
	BaseEntity
}
