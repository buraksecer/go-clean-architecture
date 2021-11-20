package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	IssuedAt   int64  `json:"iat"`
	Expiration int64  `json:"exp"`
	UserName   string `json:"username"`
}

func ValidateToken(t string, secret string) (c Claims, err error) {
	token, err := jwt.ParseWithClaims(
		t,
		&c,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil || !token.Valid {
		return c, fmt.Errorf("invalid token")
	}
	return c, nil
}

func (c *Claims) Valid() error {
	if c.Expiration == 0 || c.UserName == "" {
		return fmt.Errorf("missing jwt")
	}

	now := time.Now()
	exp := time.Unix(c.Expiration, 0)
	if now.After(exp) {
		return fmt.Errorf("token is expired")
	}
	return nil
}
