package model

import (
	"gorm.io/gorm"
)

type StudentExam struct {
	gorm.Model
	StudentID int
	Student   Student
	ExamID    int
	Exam      Exam
	RoomID    string
	Room      Room
	Grade     int `json:"grade"`
}
