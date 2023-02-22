package utils

import (
	"fiber-base-go/internal/model"
	"testing"
	"time"
)

func TestShuffleStudents(t *testing.T) {
	// Create a slice of students to shuffle
	students := []*model.Student{
		{
			Name:     "minh",
			Class:    "A1",
			Birthday: time.Now(),
		},
		{
			Name:     "minh 1",
			Class:    "A2",
			Birthday: time.Now(),
		},
		{
			Name:     "minh 2",
			Class:    "A3",
			Birthday: time.Now(),
		},
		{
			Name:     "minh 3",
			Class:    "A4",
			Birthday: time.Now(),
		},
	}

	// Shuffle the students
	ShuffleStudents(students)

	// Check that the students have been shuffled
	if students[0].Name == "Alice" && students[1].Name == "Bob" && students[2].Name == "Charlie" && students[3].Name == "Dave" && students[4].Name == "Eve" {
		t.Error("Students were not shuffled")
	}
}
