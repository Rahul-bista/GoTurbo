package goturbo

import (
	"net/http"
)

type Server struct {
	router *Router
}

func NewServer() *Server {
	return &Server{router: NewRouter()}
}

func (s *Server) Use(mw MiddlewareFunc) {
	s.router.Use(mw)
}

func (s *Server) Handle(method, pattern string, handler HandlerFunc) {
	s.router.Handle(method, pattern, handler)
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
