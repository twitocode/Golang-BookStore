package books

import (
	"github.com/google/uuid"
	"github.com/twitocode/go-web/internal/database"
	"gorm.io/gorm"
)

type BooksRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BooksRepository {
	db.AutoMigrate(&database.Book{})
	return &BooksRepository{DB: db}
}

func (r *BooksRepository) GetAllBooks() ([]database.Book, error) {
	var books []database.Book
	res := r.DB.Find(&books)

	return books, res.Error
}

func (r *BooksRepository) CreateBook(data CreatBookRequest) error {
	return r.DB.Create(&database.Book{
		ID:          uuid.NewString(),
		Rating:      data.Rating,
		Title:       data.Title,
		AuthorID:    data.AuthorID,
		Description: data.Description,
	}).Error
}

func (r *BooksRepository) GetBooksAuthor(id string) (database.Author, error) {
	book := database.Book{ID: id}

	res := r.DB.Model(&database.Book{}).Joins("Author").First(&book)
	return book.Author, res.Error
}
