package inmem

import (
	"e-learning/app/models"

	"golang.org/x/crypto/bcrypt"
)

// UserInMem .
type UserInMem struct {
	Users []*models.User
}

// Create . 
func (u *UserInMem) Create(Username, Password, Email string) (int64, error){
	return 0, nil
}
// IsValid .
func (u *UserInMem) IsValid(Email, Password string) bool {
	for _, v := range u.Users {
		if v.Email == Email {
			err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(Password))
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
