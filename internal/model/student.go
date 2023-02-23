package model

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name      string    `json:"name"`
	Class     string    `json:"class"`
	Birthday  time.Time `json:"birthday"`
	ExamGroup string    `json:"exam_group"`
}
