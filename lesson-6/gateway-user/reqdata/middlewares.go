package reqdata

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type middlewareFunc func(next http.HandlerFunc) http.HandlerFunc

type Adapter struct {
	middlewares []middlewareFunc
}

func (a *Adapter) Use(next http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range a.middlewares {
		next = middleware(next)
	}
	return next
}

func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get(RequestIDHeader)
		if rid == "" {
			rid = uuid.NewV4().String()
		}
		next(w, r.WithContext(context.WithValue(r.Context(), RequestIDContextKey, rid)))
	}
}
