package models

import "time"

// Material ...
type Material struct {
	MaterialID       int
	Name             string
	Description      string
	Category         string
	ImgURL           string
	URL              string
	Author           string
	Active           bool
	UploadedByUserID int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
