package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nchdatta/go-note-app/helpers"
)

func AllNotes(c *fiber.Ctx) error {
	return c.JSON(helpers.NewResponse(true, "Got All Notes!", nil))
}
func GetNote(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(helpers.NewResponse(true, "Got Note!"+id, nil))
}
func CreateNote(c *fiber.Ctx) error {
	return c.JSON(helpers.NewResponse(true, "Note Created!", nil))
}
func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(helpers.NewResponse(true, "Note Updated!"+id, nil))
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(helpers.NewResponse(true, "Note Deleted!"+id, nil))
}
