package middleware

import (
	"comm/pkg/services/auth"
	"context"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// check if a valid JWT is provided as bearer token. Note that the request is
// not terminated in case of errors
func ValidateBearerToken(jwtSecret string) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			userId, _ := auth.ValidateBearerToken(jwtSecret, r)
			ctx := context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
