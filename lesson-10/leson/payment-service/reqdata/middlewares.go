package reqdata

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get(RequestIDHeader)
		if rid == "" {
			rid = uuid.NewV4().String()
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), RequestIDContextKey, rid)))
	})
}
