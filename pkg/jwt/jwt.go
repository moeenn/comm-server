package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	errInvalid = errors.New("invalid or expired JWT")
)

func ValidateClientToken(secret string, token string) (string, error) {
	// the token will not parsed if it has already expired
	// expiry is checked automatically using the 'exp' claim
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !parsed.Valid {
		return "", errInvalid
	}

	// subject is the id of user to whom token was issued
	subject, err := parsed.Claims.GetSubject()
	if err != nil || subject == "" {
		return subject, errInvalid
	}

	return subject, nil
}

func ValidateServerToken(secret string, token string) error {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !parsed.Valid {
		return err
	}

	return nil
}
