package main

import (
	"shoplist/pkg/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	e.GET("/items", handlers.GetAllHandler(App.DataStore)) //return all list items
	e.GET("/items/:id", nil)                               //return item by id
	e.POST("/items", handlers.PostHandler(App.DataStore))  //insert an item
	e.DELETE("/items/:id", nil)                            //delete an item
	e.PUT("/items/:id", nil)                               //update an item
}
