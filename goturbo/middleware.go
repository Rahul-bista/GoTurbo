package goturbo

import (
	"net/http"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func (r *Router) Use(mw MiddlewareFunc) {
	for pattern, handler := range r.handlers {
		r.handlers[pattern] = mw(handler)
	}
}
