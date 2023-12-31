package auth

import (
	"comm/pkg/jwt"
	"net/http"
)

func ValidateQueryToken(jwtSecret string, r *http.Request) (string, error) {
	token, err := GetTokenFromQuery(r)
	if err != nil {
		return "", err
	}

	userId, err := jwt.ValidateClientToken(jwtSecret, token)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func ValidateBearerToken(jwtSecret string, r *http.Request) error {
	token, err := GetTokenFromAuthHeader(r)
	if err != nil {
		return err
	}

	if err := jwt.ValidateServerToken(jwtSecret, token); err != nil {
		return err
	}

	return nil
}
