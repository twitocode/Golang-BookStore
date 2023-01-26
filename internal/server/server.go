package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/twitocode/go-web/internal/authors"
	"github.com/twitocode/go-web/internal/books"
	"gorm.io/gorm"
)

type Server struct {
	Port   string
	DB     *gorm.DB
	Router chi.Router
}

func NewServer(port string, db *gorm.DB) *Server {
	return &Server{
		Port:   port,
		Router: chi.NewRouter(),
		DB:     db,
	}
}

func (s *Server) Run() error {
	br := books.NewBookRepository(s.DB)
	ar := authors.NewAuthorRepository(s.DB)

	bc := books.NewBooksController(br, ar)
	ac := authors.NewAuthorsController(ar)

	s.Router.Route(bc.Route, bc.RegisterRoutes)
	s.Router.Route(ac.Route, ac.RegisterRoutes)

	return http.ListenAndServe(s.Port, s.Router)
}
