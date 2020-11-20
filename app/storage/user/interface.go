package user

// User .
type User interface {
	// List() []*models.User
	// Get(UserID int) *models.User
	IsValid(Email, Password string) bool
	// IsVerified(UserID int) bool
	// IsActive(UserID int) bool
	Create(Username, Password, Email string) (int64, error)
	// Update(Password, Email string) (int64, error)
}
