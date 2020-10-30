package models

import (
	"time"

	"github.com/revel/revel"
)

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

// Validate .
func (user *User) Validate(v *revel.Validation) {
	v.Required(user.Email)
	v.Required(user.Password)
}
