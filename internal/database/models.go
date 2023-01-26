package database

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Books []Book

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Book struct {
	ID          string  `json:"id"`
	Rating      float32 `json:"rating"`
	Title       string  `json:"title" gorm:"unique"`
	Description string  `json:"description"`

	Author   Author `json:"author"`
	AuthorID string `json:"authorId"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
