package handlers

import (
	"net/http"
	"shoplist/pkg/models"
	"shoplist/pkg/storage"
	"strconv"

	"github.com/labstack/echo/v4"
)

type jsonResponse struct {
	Message string            `json:"message"`
	Content []models.ListItem `json:"content"`
}

func GetAllHandler(db storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid, err := strconv.Atoi(c.FormValue("user-id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "invalid uid provided",
				Content: nil,
			})
		}
		items, err := db.GetAll(uid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: err.Error(),
				Content: nil,
			})
		}
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
		uid, err := strconv.Atoi(c.FormValue("user-id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "invalid uid provided",
				Content: nil,
			})
		}
		id, err := db.Insert(title, desc, uid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: err.Error(),
				Content: nil,
			})
		}
		jsonPayload := jsonResponse{
			Message: "successful",
			Content: []models.ListItem{
				{
					Id:          id,
					Title:       title,
					Description: desc,
					UserId:      uid,
				},
			},
		}
		return c.JSON(http.StatusOK, jsonPayload)
	}
}
