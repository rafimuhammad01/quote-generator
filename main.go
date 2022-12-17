package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rafimuhammad01/quote-generator/database"
	"rafimuhammad01/quote-generator/internal"

	"github.com/gin-contrib/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	dbInit *sqlx.DB

	repoInit    internal.Repository
	serviceInit internal.Service
	handlerInit *internal.Handler
)

func InitDependency() {
	dbInit = database.PostgresInit(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("ENV"),
	)

	repoInit = internal.NewRepository(dbInit)
	serviceInit = internal.NewService(repoInit)
	handlerInit = internal.NewHandler(serviceInit)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found")
	}

	InitDependency()

	app := fiber.New(fiber.Config{ErrorHandler: internal.ErrorHandler})
	app.Use(cors.New())

	api := app.Group("/api/v1")
	api.Get("/generate-quote", handlerInit.GenerateQuote)
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "resource not found",
		})
	})

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("error when starting server: %s", err.Error())
	}
}
