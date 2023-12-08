package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Declering all the routes
func SetUpRoutes(app *fiber.App) {
	api := app.Group("api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Note Routes
	note := api.Group("note")
	note.Get("/", AllNotes)
	note.Get("/:id", GetNote)
	note.Post("/create", CreateNote)
	note.Put("/update/:id", UpdateNote)
	note.Delete("/delete/:id", DeleteNote)
}
