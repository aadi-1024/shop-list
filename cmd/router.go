package main

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"shoplist/pkg/handlers"
	"time"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	needAuth := e.Group("/items", VerifyToken)
	needAuth.GET("/", handlers.GetAllHandler(App.DataStore)) //return all list items
	needAuth.GET("/:id", nil)                                //return item by id
	needAuth.POST("/", handlers.PostHandler(App.DataStore))  //insert an item
	needAuth.DELETE("/:id", nil)                             //delete an item
	needAuth.PUT("/:id", nil)                                //update an item
	e.POST("/login", func(c echo.Context) error {
		user := c.FormValue("username")
		pass := c.FormValue("password")
		type jsonRes struct {
			Content string
			Message string
		}
		if user == "user" && pass == "password" {
			cl := &Claims{
				1,
				jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
				},
			}
			s, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, cl).SignedString([]byte("big_secret"))
			return c.JSON(200, jsonRes{
				Content: s,
				Message: "authenticated",
			})
		}
		return c.JSON(http.StatusBadRequest, jsonRes{
			Content: "",
			Message: "not authorized",
		})
	})
}
