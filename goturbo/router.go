package goturbo

import (
	"net/http"
)

type Router struct {
	handlers map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) handleRequest(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}
