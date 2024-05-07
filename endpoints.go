package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *DB
}

func setupRoutes(e *echo.Echo, db *DB) {
	handler := &Handler{DB: db}
	e.GET("/users/:id", handler.getUser)
}

func (h *Handler) getUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.DB.GetUser(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	return c.JSON(http.StatusOK, user)
}