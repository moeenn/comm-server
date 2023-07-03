package service

import (
	"comm/pkg/jwt"
	"errors"
	"net/http"
)

func ValidateToken(jwtSecret string, r *http.Request) (string, error) {
	token := r.URL.Query().Get("token")
	if token == "" {
		return "", errors.New("please provide a valid jwt as query parameter ('token')")
	}

	userId, err := jwt.ValidateToken(jwtSecret, token)
	if err != nil {
		return "", err
	}

	return userId, nil
}
