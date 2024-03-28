package main

import (
	"crypto/rand"
	"log"
	"shoplist/pkg/postgres"

	"github.com/labstack/echo/v4"
)

var App Config

func main() {
	db, err := postgres.NewDb("postgres://postgres:password@localhost:5432/list")
	if err != nil {
		log.Fatal(err)
	}
	App.DataStore = db

	e := echo.New()
	e.Debug = true
	e.HideBanner = true

	if e.Debug {
		App.jwt_secret = []byte("secret")
	} else {
		secret := make([]byte, 64)
		_, err = rand.Read(secret)
		if err != nil {
			log.Fatal(err)
		}
		App.jwt_secret = secret
	}
	SetupRouter(e)

	if err := e.Start("localhost:8080"); err != nil {
		log.Println(err)
	}
}
