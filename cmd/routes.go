package main

import (
	"github.com/WBGreenArrow/go-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFact)
	app.Post("/fact", handlers.CreateFact)
}