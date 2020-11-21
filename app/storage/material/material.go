package material

import (
	"e-learning/app/models"
)

// Repository .
type Repository struct {
	service Material
}

// List .
func (r *Repository) List() ([]*models.Material, error) {
	return r.service.List()
}

// Get .
func (r *Repository) Get(MaterialID int) (*models.Material, error) {
	return r.service.Get(MaterialID)
}

// Create .
func (r *Repository) Create(Name, Description, Category, ImgURL, URL, Author string, UploadedByUser int) (int64, error) {
	return r.service.Create(Name, Description, Category, ImgURL, URL, Author, UploadedByUser)
}

// Update .
func (r *Repository) Update(MaterialID int, Name, Description, Category, ImgURL, URL, Author string) (int64, error) {
	return r.service.Update(MaterialID, Name, Description, Category, ImgURL, URL, Author)
}

// Delete .
func (r *Repository) Delete(MaterialID int) (int64, error) {
	return r.service.Delete(MaterialID)
}

// NewMaterial .
func NewMaterial(m Material) *Repository {
	return &Repository{
		service: m,
	}
}
