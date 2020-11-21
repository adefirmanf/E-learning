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
		if material.Active{
			materials = append(materials, &material)
		}
	}
	return materials, nil 
}

// Create . 
func (m *Material) Create(Name, Description, Category, ImgURL, URL, Author string, UploadedByUser int) (int64, error) {
	q := fmt.Sprintf(`
	INSERT INTO "material" 
	("name", "description", "category", "img_url", "url", "author", "active", "upload_by_user", "created_at", "updated_at") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)

	stmt, err := m.db.Prepare(q)
	if err != nil {
		return 0, err
	}
	currentTime := time.Now()
	results, err := stmt.Exec(Name, Description, Category, ImgURL, URL, Author, true, UploadedByUser, currentTime, currentTime)
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
func (m *Material) Update(MaterialID int, Name, Description, Category, ImgURL, URL, Author string) (int64, error){
	q := fmt.Sprintf(`UPDATE "material" SET "name" = $1, "description" = $2, "category" = $3, "img_url" = $4, "url" = $5, "author" = $6, "updated_at" = $7 WHERE "id" = $8`)
	stmt, err := m.db.Prepare(q)
	if err != nil {
		return 0, err
	}
	currentTime := time.Now()
	results, err := stmt.Exec(Name, Description, Category, ImgURL, URL, Author, currentTime, MaterialID)
	
	return results.RowsAffected() 
}

// Delete . 
func (m *Material) Delete(MaterialID int) (int64, error){
	q := fmt.Sprintf(`UPDATE "material" SET "active" = $1, "updated_at" = $2 WHERE "id" = $3`)
	stmt, err := m.db.Prepare(q)
	if err != nil {
		return 0, err
	}
	currentTime := time.Now()
	results, err := stmt.Exec(false, currentTime, MaterialID)

	println(MaterialID)
	return results.RowsAffected()
}
// NewMaterialDB . 
func NewMaterialDB(sql *sql.DB) *Material{
	return &Material{
		db : sql, 
	}
}