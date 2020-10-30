package material

import (
	"e-learning/app/models"
)

// Material .
type Material interface {
	List() ([]*models.Material, error)
	Get(MaterialID int) (*models.Material, error)
	Create(Name, Description, Category, ImgURL, URL, Author string, UploadByUserID int) (int64, error)
	Update(MaterialID int, Name, Description, Category, ImgURL, URL, Author string, UploadByUserID int) (int64, error)
	IsActive(MaterialID int) bool
}
