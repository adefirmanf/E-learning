package db

import (
	"e-learning/app/models"
	"fmt"
	"time"
	"database/sql"
)

// Material . 
type Material struct{
	db *sql.DB
}

// List . 
func (m *Material) List() ([]*models.Material, error){
	var materials []*models.Material
	q := fmt.Sprintf(`SELECT "id", "name", "description", "category", "img_url", "url", "author", "active", "upload_by_user", "created_at", "updated_at" FROM "material"`)
	rows, err := m.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var material models.Material
		err := rows.Scan(&material.MaterialID, &material.Name, &material.Description, &material.Category, &material.ImgURL, &material.URL, &material.Author, &material.Active, &material.UploadedByUserID, &material.CreatedAt, &material.UpdatedAt)
		if err != nil {
			return nil, err
		}
		materials = append(materials, &material)
	}
	return materials, nil 
}

// Create . 
func (m *Material) Create(Name, Description, Category, ImgURL, URL, Author string, UploadedByUser int) (int64, error) {
	q := fmt.Sprintf(`INSERT INTO "material" ("name", description", "category", "img_url", "url", "author", "active", "upload_by_user" "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	stmt, err := m.db.Prepare(q)
	if err != nil {
		return 0, err
	}
	currentTime := time.Now()
	results, err := stmt.Exec(Name, Description, Category, ImgURL, URL, Author, UploadedByUser, currentTime, currentTime)
	if err != nil {
		return 0, err
	}
	return results.RowsAffected() 
}

// Get . 
func (m *Material) 	Get(MaterialID int) (*models.Material, error){
	var material models.Material
	q := fmt.Sprintf(`SELECT "id", "name", "description", "category", "img_url", "url", "author", "active", "upload_by_user", "created_at", "updated_at" FROM "material" WHERE "id" = $1`)
	err := m.db.QueryRow(q, MaterialID).Scan(
		&material.MaterialID, 
		&material.Name, 
		&material.Description, 
		&material.Category,
		&material.ImgURL, 
		&material.URL, 
		&material.Author,
		&material.Active,
		&material.UploadedByUserID,
		&material.CreatedAt,
		&material.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

// IsActive . 
func (m *Material) IsActive(MaterialID int) bool {
	return true
}
// Update . 
func (m *Material) Update(MaterialID int, Name, Description, Category, ImgURL, URL, Author string, UploadByUserID int) (int64, error){
	return 0, nil 
}

// NewMaterialDB . 
func NewMaterialDB(sql *sql.DB) *Material{
	return &Material{
		db : sql, 
	}
}