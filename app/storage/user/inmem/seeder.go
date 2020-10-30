package inmem

import (
	"e-learning/app/models"
)

// SeederUser .
func SeederUser() []*models.User {
	return []*models.User{
		&models.User{
			ID:       1,
			Username: "admin",
			Email:    "admin@mail.com",
			Password: "$2a$14$1we2T9X74ZlDnzzbQJwaKeZzs6O1Tr3b2ba1obhSvPgVffCHwZ.W6",
		},
	}
}
