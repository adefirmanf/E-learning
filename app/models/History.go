package models

import (
	"time"
)

// History .
type History struct {
	ID         int
	UserID     int
	MaterialID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
