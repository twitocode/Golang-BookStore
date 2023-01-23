package books

import (
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

func (r BooksRepository) GetAllBooks() []database.Book{
  var books []database.Book
  r.DB.Find(&books)

  return books
}

func (r BooksRepository) CreateBook(book *database.Book) {
  r.DB.Create(&book)
}