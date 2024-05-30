package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *DB
}

func setupRoutes(e *echo.Echo, db *DB) {
	handler := &Handler{DB: db}
	e.GET("/users/:id", handler.getUser)
	e.GET("/tournaments", handler.getTournaments)
	e.GET("/tournaments/:id", handler.getTournament)
	e.GET("/match/:id", handler.getMatch)
	e.POST("/match/:id/bet", handler.postBet)
}

func (h *Handler) getUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.DB.GetUser(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) getTournament(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	tournament, err := h.DB.GetTournament(idInt)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	return c.JSON(http.StatusOK, tournament)
}

func (h *Handler) getTournaments(c echo.Context) error {
	tournaments, err := h.DB.GetTournaments()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	return c.JSON(http.StatusOK, tournaments)
}

func (h *Handler) getTournamentBets(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	bets, err := h.DB.GetTournamentBets(idInt)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	return c.JSON(http.StatusOK, bets)
}

func (h *Handler) getMatch(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	match, err := h.DB.GetMatch(idInt)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, match)
}

func (h *Handler) postBet(c echo.Context) error {
	resBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	var bet Bet
	if err := json.Unmarshal(resBody, &bet); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if bet.MatchID <= 0 && (bet.Bet == "HOME" || bet.Bet == "AWAY" || bet.Bet == "DRAW")  {
		return c.NoContent(http.StatusBadRequest)
	}

	err = h.DB.PostBet(bet)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// ?
	return c.NoContent(http.StatusOK)
}