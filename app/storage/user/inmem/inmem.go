package inmem

import (
	"e-learning/app/models"

	"golang.org/x/crypto/bcrypt"
)

// UserInMem .
type UserInMem struct {
	Users []*models.User
}

// IsValid .
func (u *UserInMem) IsValid(email, password string) bool {
	for _, v := range u.Users {
		if v.Email == email {
			err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(password))
			if err != nil {
				return false
			}
			return true
		}
	}
	return false
}

// NewUserInMem .
func NewUserInMem(u []*models.User) *UserInMem {
	return &UserInMem{
		Users: u,
	}
}
