package authors

import (
	"github.com/google/uuid"
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

func (r AuthorsRepository) GetAllAuthors() ([]database.Author, error) {
	var authors []database.Author
	res := r.DB.Find(&authors)

	return authors, res.Error
}

func (r AuthorsRepository) GetAuthorById(id string) (database.Author, error) {
	var author database.Author
	res := r.DB.Find(&author, database.Author{ID: id})

	return author, res.Error
}

func (r AuthorsRepository) CreateAuthor(data CreateAuthorRequest) error {
	return r.DB.Create(&database.Author{
		ID:        uuid.NewString(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}).Error
}
