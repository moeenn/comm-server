package auth

import (
	"comm/pkg/jwt"
	"comm/pkg/services/token"
	"net/http"
)

func ValidateQueryToken(jwtSecret string, r *http.Request) (string, error) {
	token, err := token.GetTokenFromQuery(r)
	if err != nil {
		return "", err
	}

	userId, err := jwt.ValidateToken(jwtSecret, token)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func ValidateBearerToken(jwtSecret string, r *http.Request) (string, error) {
	token, err := token.GetTokenFromAuthHeader(r)
	if err != nil {
		return "", err
	}

	userId, err := jwt.ValidateToken(jwtSecret, token)
	if err != nil {
		return "", err
	}

	return userId, nil
}
