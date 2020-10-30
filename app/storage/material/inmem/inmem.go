package inmem

import (
	"e-learning/app/models"
)

// MaterialInMem .
type MaterialInMem struct {
	Materials []*models.Material
}

// List .
func (m *MaterialInMem) List() ([]*models.Material, error) {
	return m.Materials, nil
}

// Get .
func (m *MaterialInMem) Get(MaterialD int) (*models.Material, error) {
	var material *models.Material
	for _, v := range m.Materials {
		if v.MaterialID == MaterialD {
			material = v
		}
	}
	return material, nil
}

// Create .
func (m *MaterialInMem) Create(Name, Description, Category, ImgURL, URL, Author string, UplaodedByUserID int) (int64, error) {
	return 0, nil
}

// Update .
func (m *MaterialInMem) Update(MaterialID int, Name, Description, Category, ImgURL, URL, Author string, UplaodedByUserID int) (int64, error) {
	return 0, nil
}

// IsActive .
func (m *MaterialInMem) IsActive(MaterialID int) bool {
	for _, v := range m.Materials {
		if (v.MaterialID == MaterialID) && v.Active == true {
			return true
		}
	}
	return false
}

// NewMaterialInMem .
func NewMaterialInMem(m []*models.Material) *MaterialInMem {
	return &MaterialInMem{
		Materials: m,
	}
}
