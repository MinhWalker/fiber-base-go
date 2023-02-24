package services

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"

	"fiber-base-go/internal/model"
)

type PopulateService interface {
	Populate() error
}

type populateService struct {
	roomService    RoomService
	studentService StudentService
}

var _ PopulateService = (*roomService)(nil)

func NewPopulateService(roomService RoomService, studentService StudentService) PopulateService {
	return &populateService{roomService, studentService}
}

func (s *populateService) Populate() error {
	if err := s.PopulateRoom(); err != nil {
		log.Fatalf("PopulateRoom failed: %s\n", err)
		return err
	}
	if err := s.PopulateStudent(); err != nil {
		log.Fatalf("PopulateStudents failed: %s\n", err)
		return err
	}
	return nil
}

func (s *populateService) PopulateRoom() error {
	file, err := os.Open("sample_data/rooms.csv")
	if err != nil {
		return errors.Wrap(err, "populateService.PopulateRoom")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return errors.Wrap(err, "populateService.PopulateRoom")
	}
	rooms := make([]*model.Room, 0)
	for _, record := range records[1:] {
		room := &model.Room{
			Name:     record[0],
			Capacity: record[1],
		}
		rooms = append(rooms, room)
	}
	return s.roomService.importRooms(rooms)
}

func (s *populateService) PopulateStudent() error {
	file, err := os.Open("sample_data/students.csv")
	if err != nil {
		return errors.Wrap(err, "populateService.PopulateStudent")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return errors.Wrap(err, "populateService.PopulateStudent")
	}
	students := make([]*model.Student, 0)
	for _, record := range records[1:] {
		birthday, _ := time.Parse("02/01/2006", record[2])
		student := &model.Student{
			Name:      record[0],
			Class:     record[1],
			Birthday:  birthday,
			ExamGroup: record[3],
		}
		students = append(students, student)
	}
	return s.studentService.ImportStudent(students)
}
