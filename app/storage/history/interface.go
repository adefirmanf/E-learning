package history

import (
	"e-learning/app/models"
)

// History .
type History interface {
	List() ([]*models.History, error)
	Get(ID int) (*models.History, error)
	Create(UserID int, MaterialID int) (int64, error)
}
