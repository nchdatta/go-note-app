package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nchdatta/go-note-app/database"
	"github.com/nchdatta/go-note-app/helpers"
	"github.com/nchdatta/go-note-app/models"
)

func AllNotes(c *fiber.Ctx) error {
	notes := []models.Note{}

	database.DBConn.Find(&notes)

	return c.JSON(helpers.NewResponse(true, "Got All Notes!", notes))
}
func GetNote(c *fiber.Ctx) error {
	id := c.Params("id")
	note := models.Note{}

	database.DBConn.First(&note, id)

	if note.Id == uuid.Nil {
		return c.Status(404).JSON(helpers.NewResponse(false, "Note not exists.", nil))
	}

	return c.JSON(helpers.NewResponse(true, "Got Note!"+id, note))
}
func CreateNote(c *fiber.Ctx) error {
	note := models.Note{}
	if err := c.BodyParser(&note); err != nil {
		return c.Status(400).JSON(helpers.NewResponse(false, err.Error(), nil))
	}
	note.Id = uuid.New()

	err := database.DBConn.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(helpers.NewResponse(false, "Could not create note.", err))
	}

	return c.JSON(helpers.NewResponse(true, "Note Created! ", note))
}
func UpdateNote(c *fiber.Ctx) error {
	// Update Struct
	UpdateNote := struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Details  string `json:"details"`
	}{}

	id := c.Params("id")
	note := models.Note{}
	database.DBConn.First(&note, "id = ?", id)

	if note.Id == uuid.Nil {
		return c.Status(404).JSON(helpers.NewResponse(false, "Note not exists.", nil))
	}

	if err := c.BodyParser(&UpdateNote); err != nil {
		return c.Status(400).JSON(helpers.NewResponse(false, err.Error(), nil))
	}

	note.Title = UpdateNote.Title
	note.Title = UpdateNote.SubTitle
	note.Title = UpdateNote.Details

	err := database.DBConn.Where("id = ?", id).UpdateColumns(&note).Error
	if err != nil {
		return c.Status(500).JSON(helpers.NewResponse(false, "Could not update note.", err))
	}

	return c.JSON(helpers.NewResponse(true, "Note Updated!"+id, nil))
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	note := models.Note{}

	database.DBConn.First(&note, "id = ?", id)
	if note.Id == uuid.Nil {
		return c.Status(404).JSON(helpers.NewResponse(false, "Note not exists.", nil))
	}

	err := database.DBConn.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(500).JSON(helpers.NewResponse(false, "Could not delete note.", err))
	}

	return c.JSON(helpers.NewResponse(true, "Note Deleted!"+id, nil))
}
