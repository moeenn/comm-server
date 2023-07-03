package token

import (
	"errors"
	"net/http"
	"strings"
)

func GetTokenFromQuery(r *http.Request) (string, error) {
	token := r.URL.Query().Get("token")
	if token == "" {
		return "", errors.New("please provide a valid jwt as query parameter ('token')")
	}

	return token, nil
}

func GetTokenFromAuthHeader(r *http.Request) (string, error) {
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
