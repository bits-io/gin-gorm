package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}