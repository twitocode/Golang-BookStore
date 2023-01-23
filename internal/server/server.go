package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/twitocode/go-web/internal/books"
	"github.com/twitocode/go-web/internal/authors"
	"gorm.io/gorm"
)

type Server struct {
	Port   string 
  DB *gorm.DB
	Router chi.Router
}

func NewServer(port string, db *gorm.DB) *Server {
	return &Server{
		Port:   port,
		Router: chi.NewRouter(),
    DB: db,
	}
}

func (s *Server) Run() error {
  br := books.NewBookRepository(s.DB)
  ar := authors.NewAuthorRepository(s.DB)
    
  bh := books.NewBooksHandler(br, ar)
  ah := authors.NewAuthorsHandler(ar)


	s.Router.Route(bh.Route, bh.RegisterRoutes)
	s.Router.Route(ah.Route, ah.RegisterRoutes)

	return http.ListenAndServe(s.Port, s.Router)
}
