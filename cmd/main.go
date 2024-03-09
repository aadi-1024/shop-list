package main

import (
	"shoplist/pkg/memstore"

	"github.com/labstack/echo/v4"
)

var App Config

func main() {
	App.DataStore = memstore.NewMemStore()
	e := echo.New()
	e.Debug = true
	SetupRouter(e)

	e.Start("localhost:8080")
}
