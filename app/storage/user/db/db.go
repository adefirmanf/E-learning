package db

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"e-learning/app/models"
	"fmt"
	"database/sql"
)

// User . 
type User struct{
	db *sql.DB
}

// IsValid . 
func (u *User) IsValid(Email, Password string) bool {
	var user models.User
	q := fmt.Sprintf(`SELECT "id", "username", "password", "email", "is_verified", "role" FROM "user" WHERE "email" = $1`)
	err := u.db.QueryRow(q, Email).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.IsVerified, &user.Role)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if err != nil {
		return false
	}
	return user.IsVerified
}

// Create . 
func (u *User) Create(Username, Password, Email string) (int64, error){
	q := fmt.Sprintf(`INSERT INTO "user" ("username", "password", "email", "is_verified", "role", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	stmt, err := u.db.Prepare(q)
	if err != nil {
		return 0, err 
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.MinCost)
	if err != nil {
		return 0, err
	}
	currentTime := time.Now()
	results, err := stmt.Exec(Username, bytes, Email, true, "member", currentTime, currentTime)
	if err != nil {
		return 0, err
	}
	return results.RowsAffected()
}

// NewUserDB . 
func NewUserDB(sql *sql.DB) *User{
	return &User{
		db : sql,
	}
}