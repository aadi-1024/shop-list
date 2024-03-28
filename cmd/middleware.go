package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Claims struct {
	Userid int `json:"uid"`
	jwt.RegisteredClaims
}

type jsonRes struct {
	Message string `json:"message"`
	Content string `json:"content,omitempty"`
}

// middleware wrapper
func VerifyToken(jwtToken []byte) func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	//middleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, jsonRes{
					Message: "Unauthorized",
					Content: "",
				})
			}
			temp := strings.Split(authHeader, " ")
			if len(temp) != 2 {
				return c.JSON(http.StatusUnauthorized, jsonRes{
					Message: "No token detected",
					Content: "",
				})
			}
			bearerToken := temp[1]
			clm := &Claims{}
			token, err := jwt.ParseWithClaims(bearerToken, clm, func(token *jwt.Token) (interface{}, error) {
				return jwtToken, nil
			})

			if errors.Is(err, jwt.ErrTokenMalformed) {
				return c.JSON(http.StatusInternalServerError, jsonRes{
					Message: "Token malformed",
					Content: "",
				})
			}
			if !token.Valid {
				return c.JSON(http.StatusUnauthorized, jsonRes{
					Message: "Auth token could not be validated",
					Content: "",
				})
			}
			if clm.ExpiresAt.Before(time.Now()) {
				return c.JSON(http.StatusUnauthorized, jsonRes{
					Message: "Session expired. Log in again",
					Content: "",
				})
			}
			c.Request().Header.Add("user-id", strconv.Itoa(clm.Userid))
			return next(c)
		}
	}
}
