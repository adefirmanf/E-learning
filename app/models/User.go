package models

import "time"

// User .
type User struct {
	ID         int
	Username   string
	Password   string
	Email      string
	Role       string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
