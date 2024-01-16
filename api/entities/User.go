package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"uniqueIndex;type:VARCHAR(50)"`
	Email    string  `json:"email" gorm:"uniqueIndex;type:VARCHAR(255)"`
	Password string  `json:"password"`
	Tweets   []Tweet `gorm:"foreignKey:UserID"`
}
