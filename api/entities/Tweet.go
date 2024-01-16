package entities

import (
	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	User        User   `json:"user"`
	Description string `json:"description"`
}
