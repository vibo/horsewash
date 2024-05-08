package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/idtoken"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	clientID := os.Getenv("CLIENT_ID")
	dbURL := os.Getenv("DB_URL")

	db, err := openDatabase(dbURL)
	if err != nil {
		log.Printf("Error: %v", err)
		log.Fatal("Unable to open a connection to the database")
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			idToken := strings.TrimPrefix(authorization, "Bearer ")

			payload, err := idtoken.Validate(c.Request().Context(), idToken, clientID)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to verify ID token")
			}

			if email, ok := payload.Claims["email"].(string); ok {
				fmt.Println(email)
			}

			c.Set("user_id", payload.Subject)
			return next(c)
		} 
	})
	
	e.Use(middleware.Logger())

	setupRoutes(e, db)

	e.Start(":8080")
}