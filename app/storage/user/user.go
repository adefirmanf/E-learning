package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Repository .
type Repository struct {
	Service User
}

// List .
// func (r *Repository) List() []*models.User {
// 	return r.Service.List()
// }

// IsValid .
func (r *Repository) IsValid(Email, Password string) bool {
	return r.Service.IsValid(Email, Password)
}

// HashPassword .
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// CheckPasswordAndHash .
func checkPasswordAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NewUser .
func NewUser(u User) *Repository {
	return &Repository{
		Service: u,
	}
}
