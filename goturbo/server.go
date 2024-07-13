package framework

import (
	"net/http"
)

type Server struct {
	router *Router
}

func NewServer() *Server {
	return &Server{
		router: NewRouter(),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.handleRequest(w, r)
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}
