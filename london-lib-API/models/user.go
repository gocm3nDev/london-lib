package models

import "time"

type User struct {
	Name         string    `json:"Name"`
	Email        string    `json:"Email"`
	PasswordHash string    `json:"PasswordHash"`
	IsActive     bool      `json:"IsActive"`
	CreatedAt    time.Time `json:"CreatedAt"`
}
