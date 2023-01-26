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

type BooksController struct {
	Repository       *BooksRepository
	AuthorRepository *authors.AuthorsRepository
	Route            string
}

func NewBooksController(repo *BooksRepository, aRepo *authors.AuthorsRepository) *BooksController {
	return &BooksController{
		Route:            "/books",
		Repository:       repo,
		AuthorRepository: aRepo,
	}
}

func (h *BooksController) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAllBooks)
	r.Get("/author/{id}", h.GetBooksAuthor)
	r.Post("/", h.CreateBook)
	r.Patch("/{id}", h.UpdateBook)
	r.Delete("/", h.DeleteBook)
}

func (h *BooksController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	books, err := h.Repository.GetAllBooks()

	if err != nil {
		encoder.Encode(common.ErrorResponse{Code: 500, Error: "Could not query books"})
	}

	encoder.Encode(common.Response[[]database.Book]{Code: 200, Data: books})
}

func (h *BooksController) CreateBook(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	var body CreatBookRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Fatalf("Could not read the request body")
		encoder.Encode(common.ErrorResponse{Code: 400, Error: "Could not read the request body"})
    return
	}

	err := h.Repository.CreateBook(body)

	if err != nil {
		encoder.Encode(common.ErrorResponse{Code: 500, Error: "Could not create the book"})
    return
	}

	encoder.Encode(common.Response[string]{Code: 201, Data: "Successfully created the book"})
}

func (h *BooksController) GetBooksAuthor(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")
	author, err := h.Repository.GetBooksAuthor(bookId)
	encoder := json.NewEncoder(w)

	if err != nil {
		encoder.Encode(common.ErrorResponse{Code: 500, Error: "Could not get the author"})
    return
	}

	json.NewEncoder(w).Encode(common.Response[database.Author]{Code: 200, Data: author})
}

func (h *BooksController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (h *BooksController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
