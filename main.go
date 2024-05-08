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
	if clientID == "" {
		log.Fatal("CLIENT_ID is not set in the environment variables")
	}

	dbUrl := os.Getenv("DB_URL")
	if clientID == "" {
		log.Fatal("DB_URL is not set in the environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := openDatabase(dbUrl)
	if err != nil {
		log.Printf("Error: %v", err)
		log.Fatal("Unable to open a connection to the database")
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

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
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Email claim is missing in ID token")
			}

			c.Set("user_id", payload.Subject)
			return next(c)
		} 
	})
	
	e.Use(middleware.Logger())

	setupRoutes(e, db)

	e.Start(":" + port)
}