package utils

import (
	"encoding/csv"
	"io"
	"time"

	"fiber-base-go/internal/model"
)

func ParseCSV(r io.Reader, batchSize int) ([][]*model.Student, error) {
	reader := csv.NewReader(r)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	header := rows[0]
	dataRows := rows[1:]

	// Map the header names to their column indexes
	var nameIndex, classIndex, birthdayIndex int
	for i, col := range header {
		switch col {
		case "name":
			nameIndex = i
		case "class":
			classIndex = i
		case "birthday":
			birthdayIndex = i
		}
	}

	students := make([][]*model.Student, 0, len(dataRows)/batchSize+1)
	batch := make([]*model.Student, 0, batchSize)

	for _, row := range dataRows {
		// Parse the birthday field into a time.Time value
		birthday, err := time.Parse("2006-01-02", row[birthdayIndex])
		if err != nil {
			return nil, err
		}

		student := &model.Student{
			Name:     row[nameIndex],
			Class:    row[classIndex],
			Birthday: birthday,
		}

		batch = append(batch, student)
		if len(batch) >= batchSize {
			students = append(students, batch)
			batch = make([]*model.Student, 0, batchSize)
		}
	}

	if len(batch) > 0 {
		students = append(students, batch)
	}

	return students, nil
}
