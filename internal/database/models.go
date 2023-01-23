package database

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Rating      float32 `json:"rating"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      Author  `json:"author"`
	AuthorId    uint    `json:"authorId"`
}

type Author struct {
	gorm.Model

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Books     []Book `json:"books"`
}
