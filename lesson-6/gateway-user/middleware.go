package main

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var RequestIDContextKey = "X-Request-ID"

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

func requestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewV4()
		rWithId := r.WithContext(context.WithValue(r.Context(), RequestIDContextKey, id.String()))
		next(w, rWithId)
	}
}
