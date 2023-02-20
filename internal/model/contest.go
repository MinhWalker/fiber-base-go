package model

import (
	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	Name     string
	semester string
	Students []*Student `gorm:"many2many:contest_students"`
}
