package authors

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/twitocode/go-web/internal/common"
	"github.com/twitocode/go-web/internal/database"
)

type AuthorsController struct {
	Repository *AuthorsRepository
	Route      string
}

func NewAuthorsController(repo *AuthorsRepository) *AuthorsController {
	return &AuthorsController{
		Route:      "/authors",
		Repository: repo,
	}
}

func (h *AuthorsController) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAllAuthors)
	r.Get("/{id}", h.GetAuthorById)
	r.Post("/", h.CreateAuthor)
}

func (h *AuthorsController) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.Repository.GetAllAuthors()
	encoder := json.NewEncoder(w)

	if err != nil {
		encoder.Encode(common.ErrorResponse{Code: 500, Error: "Could not query the authors"})
    return
	}

	encoder.Encode(authors)
}

func (h *AuthorsController) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "id")
	encoder := json.NewEncoder(w)
	author, err := h.Repository.GetAuthorById(authorId)

	if err != nil {
		encoder.Encode(common.ErrorResponse{Code: 500, Error: "Could not query the author"})
    return
	}

	encoder.Encode(common.Response[database.Author]{Code: 200, Data: author})
}

func (h *AuthorsController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	var body CreateAuthorRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Fatalf("Could not read the request body")
		encoder.Encode(common.ErrorResponse{Code: 400, Error: "Could not read the request body"})
    return
	}

	err := h.Repository.CreateAuthor(body)

	if err != nil {
		encoder.Encode(common.ErrorResponse{Code: 500, Error: "Could not create the author"})
    return
	}

	encoder.Encode(common.Response[string]{Code: 200, Data: "Successfully created the author"})
}
