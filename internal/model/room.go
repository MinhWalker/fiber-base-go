package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name     string `gorm:"unique" json:"name"`
	Capacity string `json:"capacity"`
}
