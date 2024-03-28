package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"shoplist/pkg/models"
	"shoplist/pkg/storage"
	"strconv"
	"time"
)

type jsonResponse struct {
	Message string `json:"message"`
	Content any    `json:"content,omitempty"`
}

func GetAllHandler(db storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid, err := strconv.Atoi(c.Request().Header.Get("user-id"))
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
		desc := c.FormValue("desc")
		uid, err := strconv.Atoi(c.Request().Header.Get("user-id"))
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

func GetByIdHandler(db storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		//will always be valid because of jwt
		uid, err := strconv.Atoi(c.Request().Header.Get("user-id"))
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "correct id expected",
				Content: nil,
			})
		}

		m, err := db.GetById(id, uid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "error while fetching from db",
				Content: nil,
			})
		}

		return c.JSON(http.StatusOK, jsonResponse{
			Message: "success",
			Content: m,
		})
	}
}

func PutHandler(db storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		desc := c.FormValue("desc")
		if title == "" {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "title can't be empty",
				Content: nil,
			})
		}
		//always valid because of jwt
		uid, _ := strconv.Atoi(c.Request().Header.Get("user-id"))
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "correct id expected",
				Content: nil,
			})
		}
		m := models.ListItem{
			Id:          id,
			Title:       title,
			Description: desc,
			UserId:      uid,
		}

		ret, err := db.Update(m)
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "error while fetching from db",
				Content: nil,
			})
		}

		return c.JSON(http.StatusOK, jsonResponse{
			Message: "success",
			Content: ret,
		})
	}
}

func DeleteHandler(db storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		//always valid due to JWT
		uid, _ := strconv.Atoi(c.Request().Header.Get("user-id"))
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "correct id expected",
				Content: nil,
			})
		}
		err = db.Delete(id, uid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, jsonResponse{
				Message: "error while querying db",
				Content: nil,
			})
		}
		return c.JSON(http.StatusOK, nil)
	}
}

type Claims struct {
	Userid int `json:"uid"`
	jwt.RegisteredClaims
}

func LoginPostHandler(jwtSecret []byte) echo.HandlerFunc {
	return func(c echo.Context) error {
		//sample login data, integrate with db later
		loginCreds := map[string]int{ //username userid, password is password
			"user":  1,
			"user1": 2,
		}
		username := c.FormValue("username")
		password := c.FormValue("password")

		uid, ok := loginCreds[username]
		if !ok || password != "password" {
			return c.JSON(http.StatusUnauthorized, jsonResponse{Message: "Invalid Credentials"})
		}

		clm := &Claims{
			uid,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			},
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, clm).SignedString(jwtSecret)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, jsonResponse{Message: "Something went wrong"})
		}

		return c.JSON(http.StatusOK, jsonResponse{
			Message: "Authenticated",
			Content: token,
		})
	}
}
