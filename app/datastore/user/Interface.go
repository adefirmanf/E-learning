package user

import (
	"e-learning/app/models"
)

// User .
type User interface {
	List() []*models.User
	Get(UserID int) *models.User
	IsValid(Username, Password string) bool
	IsVerified(UserID int) bool
	IsActive(UserID int) bool
	Create(Username, Password, Email string) (int64, error)
	Update(Password, Email string) (int64, error)
}
