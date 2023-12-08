package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id       uuid.UUID `json:"id" gorm:"primaryKey;"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Details  string    `json:"details"`
}
