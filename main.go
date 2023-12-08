package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/go-note-app/database"
	"github.com/nchdatta/go-note-app/routes"
)

// Initializing the database
func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	routes.SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
