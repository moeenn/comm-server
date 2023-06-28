package middleware

import (
	"comm/pkg/jwt"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ValidateToken(secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := parseToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId, err := jwt.ValidateToken(secret, token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: store userId inside request context

		fmt.Fprintf(w, "userId: %s", userId)
	}
}

func parseToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return auth, errors.New("missing Authorization header")
	}

	if !strings.Contains(auth, "Bearer ") {
		return auth, errors.New("please provide a valid bearer token")
	}

	token := strings.ReplaceAll(auth, "Bearer ", "")
	if token == "" {
		return auth, errors.New("missing token value inside Authorization header")
	}

	return token, nil
}
