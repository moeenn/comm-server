package jwt

import (
	"errors"
	"os"
	"reflect"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("failed to read JWT_SECRET value from the environment")
	}

	return secret
}

func ValidateToken(token string) (string, error) {
	secret := getJWTSecret()

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !parsed.Valid {
		return "", errors.New("invalid or expired JWT")
	}

	claims, err := extractClaims(parsed.Claims)
	if err != nil {
		return "", err
	}

	userId, ok := claims["userId"]
	if !ok {
		return "", errors.New("validated token does not contain userId")
	}

	return userId.(string), nil
}

func extractClaims(claims jwt.Claims) (map[string]any, error) {
	v := reflect.ValueOf(claims)
	result := make(map[string]any)

	if v.Kind() != reflect.Map {
		return result, errors.New("token contains unknown data-structure")
	}

	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		result[key.String()] = value.Interface()
	}

	return result, nil
}
