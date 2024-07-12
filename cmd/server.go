package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zlalabs/bunny/api"
	"github.com/zlalabs/bunny/internal/database"
)

func Execute() {
	e := echo.New()
	// e.Logger.SetLevel( slog.LevelInfo)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.New()
	api.Server(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))

		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("Server exiting")
	// endless.ListenAndServe(":8080", e)
}
