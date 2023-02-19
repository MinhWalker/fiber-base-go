package utils

import (
	"encoding/csv"
	"io"
	"time"

	"fiber-base-go/internal/model"

	"github.com/pkg/errors"
)

func ParseCSV(r io.Reader, batchSize int) ([][]*model.Student, error) {
	reader := csv.NewReader(r)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "parseCSV")
	}

	// Split the lines into batches
	batchCount := (len(lines) - 1) / batchSize
	if (len(lines)-1)%batchSize > 0 {
		batchCount++
	}
	batches := make([][]*model.Student, batchCount)
	for i := 0; i < batchCount; i++ {
		start := 1 + i*batchSize
		end := start + batchSize
		if end > len(lines) {
			end = len(lines)
		}
		batch := make([]*model.Student, 0, end-start)
		for j := start; j < end; j++ {
			line := lines[j]
			if len(line) != 3 {
				return nil, errors.New("invalid number of fields in line")
			}
			birthday, err := time.Parse("2006-01-02", line[2])
			if err != nil {
				return nil, errors.Wrap(err, "parseCSV")
			}
			student := &model.Student{
				Name:     line[0],
				Class:    line[1],
				Birthday: birthday,
			}
			batch = append(batch, student)
		}
		batches[i] = batch
	}

	return batches, nil
}
