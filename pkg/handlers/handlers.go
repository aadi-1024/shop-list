package handlers

import (
	"net/http"
	"shoplist/pkg/models"
	"shoplist/pkg/storage"

	"github.com/labstack/echo/v4"
)

type jsonResponse struct {
	Message string            `json:"message"`
	Content []models.ListItem `json:"string"`
}

func GetAllHandler(db storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		items, _ := db.GetAll()
		jsonPayload := jsonResponse{
			Message: "successful",
			Content: items,
		}
		// con, err := json.Marshal(jsonPayload)
		return c.JSON(http.StatusOK, jsonPayload)
	}
}

func PostHandler(db storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		desc := c.FormValue("description")

		id, _ := db.Insert(title, desc)
		jsonPayload := jsonResponse{
			Message: "successful",
			Content: []models.ListItem{
				{
					Id:          id,
					Title:       title,
					Description: desc,
				},
			},
		}
		return c.JSON(http.StatusOK, jsonPayload)
	}
}
