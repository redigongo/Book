package models

import (
	"gorm.io/gorm"
	"time"
)

// Book represents a book in the database.
type Book struct {
	gorm.Model
	Id             int
	Title          string
	Author         string
	BookCoverBytes []byte
	CoverFileType  string
	CoverFileName  string
	PublishedAt    time.Time
}
