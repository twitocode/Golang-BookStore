package authors

import (
	"fmt"

	"github.com/twitocode/go-web/internal/database"
	"gorm.io/gorm"
)

type AuthorsRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorsRepository {
	db.AutoMigrate(&database.Author{})
	return &AuthorsRepository{DB: db}
}

func (r AuthorsRepository) GetAllAuthors() []database.Author {
	var authors []database.Author
	r.DB.Find(&authors)

	return authors
}

func (r AuthorsRepository) GetAuthorById(id uint) database.Author {
	var author database.Author
  fmt.Println(author)
	r.DB.Find(&author, database.Author{Model: gorm.Model{ID: id}})

	return author
}

func (r AuthorsRepository) CreateAuthor(data *database.Author) {
	r.DB.Create(data)
}
