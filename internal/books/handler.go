package books

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/twitocode/go-web/internal/authors"
	"github.com/twitocode/go-web/internal/common"
	"github.com/twitocode/go-web/internal/database"
)

type BooksHandler struct {
	Repository       *BooksRepository
	AuthorRepository *authors.AuthorsRepository
	Route            string
}

func NewBooksHandler(repo *BooksRepository, aRepo *authors.AuthorsRepository) *BooksHandler {
	return &BooksHandler{
		Route:            "/books",
		Repository:       repo,
		AuthorRepository: aRepo,
	}
}

func (h *BooksHandler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAllBooks)
	r.Post("/", h.CreateBook)
	r.Patch("/{bookId}", h.UpdateBook)
	r.Delete("/", h.DeleteBook)
}

func (h *BooksHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.Repository.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func (h *BooksHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	var body database.Book

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Fatalf("Could not read the request body")
		encoder.Encode(common.Status{Code: 400, Message: "Could not read the request body"})
	}

	author := h.AuthorRepository.GetAuthorById(uint(body.AuthorId))

	h.Repository.CreateBook(&database.Book{
		Author:      author,
		AuthorId:    author.ID,
		Title:       body.Title,
		Rating:      body.Rating,
		Description: body.Description,
	})

	encoder.Encode(common.Response[string]{Code: 200, Data: "Successfully created the book"})
}

func (h *BooksHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (h *BooksHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
