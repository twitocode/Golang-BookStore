package authors

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/twitocode/go-web/internal/common"
	"github.com/twitocode/go-web/internal/database"
)

type AuthorsHandler struct {
	Repository       *AuthorsRepository
	Route            string
}

func NewAuthorsHandler(repo *AuthorsRepository) *AuthorsHandler {
	return &AuthorsHandler{
		Route:            "/authors",
		Repository:       repo,
	}
}

func (h *AuthorsHandler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAllAuthors)
	r.Get("/{authorId}", h.GetAuthorById)
  r.Post("/", h.CreateAuthor)
}

func (h *AuthorsHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	Authors := h.Repository.GetAllAuthors()
	json.NewEncoder(w).Encode(Authors)
}

func (h *AuthorsHandler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	authorId, err := strconv.ParseUint(chi.URLParam(r, "authorId"), 10, 32)
	encoder := json.NewEncoder(w)

	if err != nil {
		log.Fatalf("Not a proper author id")
		encoder.Encode(common.Status{Code: 400, Message: "Not a proper author id"})
	}

  author := h.Repository.GetAuthorById(uint(authorId))
  encoder.Encode(common.Response[database.Author]{Code:200, Data: author})
}

func (h *AuthorsHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	var body *database.Author

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Fatalf("Could not read the request body")
		encoder.Encode(common.Status{Code: 400, Message: "Could not read the request body"})
	}

	h.Repository.CreateAuthor(body)
	encoder.Encode(common.Response[string]{Code: 200, Data: "Successfully created the author"})
}
