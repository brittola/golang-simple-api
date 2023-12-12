package entities

import (
	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	User        string `json:"user"`
	Description string `json:"description"`
}
