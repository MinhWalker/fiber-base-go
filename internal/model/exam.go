package model

import (
	"time"

	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	ExamGroup string    `json:"exam_group"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	ContestID int
	Contest   Contest
}
