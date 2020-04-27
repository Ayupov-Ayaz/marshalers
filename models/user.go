package models

import "time"

//easyjson:json
type User struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Age uint8 `json:"age"`
	Email string `json:"email"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
}

func NewUser() *User {
	return &User{
		ID:          99999999999999,
		Name:        "tommy",
		Age:         28,
		Email:       "example@gmail.com",
		Created:     time.Now(),
		LastUpdated: time.Now(),
	}
}