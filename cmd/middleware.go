package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

type jwtData struct {
	secretToken []byte
	tokens      map[string]bool
}

type Claims struct {
	Userid int `json:"user-id"`
	jwt.RegisteredClaims
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	data := &jwtData{
		secretToken: []byte("big_secret"),
		tokens:      make(map[string]bool),
	}

	return func(c echo.Context) error {
		clms := &Claims{}
		hdr := c.Request().Header.Get("Authorization")
		if hdr == "" {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		bearerToken := strings.Split(hdr, " ")[1]
		_, ok := data.tokens[bearerToken]
		token, err := jwt.ParseWithClaims(bearerToken, clms, func(token *jwt.Token) (interface{}, error) {
			return data.secretToken, nil
		})
		if !ok {
			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			if token.Valid {
				data.tokens[bearerToken] = true
			}
		}
		if token.Valid {
			if clms.ExpiresAt.Before(time.Now()) {
				delete(data.tokens, bearerToken)
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			return next(c)
		}
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
