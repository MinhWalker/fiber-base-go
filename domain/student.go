package domain

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name     string    `json:"name" gorm:"text;not null;default:null`
	Class    string    `json:"class" gorm:"text;not null;default:null`
	Birthday time.Time `json:"birthday"`
}
