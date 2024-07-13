package goturbo

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type MiddlewareFunc func(HandlerFunc) HandlerFunc

type Router struct {
	routes     map[string]map[string]HandlerFunc
	middleware []MiddlewareFunc
}

func NewRouter() *Router {
	return &Router{
		routes:     make(map[string]map[string]HandlerFunc),
		middleware: []MiddlewareFunc{},
	}
}

func (r *Router) Use(mw MiddlewareFunc) {
	r.middleware = append(r.middleware, mw)
}

func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
	if r.routes[pattern] == nil {
		r.routes[pattern] = make(map[string]HandlerFunc)
	}
	// Apply middleware to the handler
	for _, mw := range r.middleware {
		handler = mw(handler)
	}
	r.routes[pattern][method] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routes[req.URL.Path]; ok {
		if handler, ok := handlers[req.Method]; ok {
			handler(w, req)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		http.NotFound(w, req)
	}
}
