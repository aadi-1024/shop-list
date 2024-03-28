package main

import (
	"github.com/labstack/echo/v4"
	"shoplist/pkg/handlers"
)

func SetupRouter(e *echo.Echo) {
	needAuth := e.Group("/items", VerifyToken(App.jwt_secret))
	needAuth.GET("/", handlers.GetAllHandler(App.DataStore))    //return all list items
	needAuth.GET("/id", handlers.GetByIdHandler(App.DataStore)) //return item by id
	needAuth.POST("/", handlers.PostHandler(App.DataStore))     //insert an item
	needAuth.DELETE("/", handlers.DeleteHandler(App.DataStore)) //delete an item
	needAuth.PUT("/", handlers.PutHandler(App.DataStore))       //update an item
	e.POST("/login", handlers.LoginPostHandler(App.jwt_secret))
}
