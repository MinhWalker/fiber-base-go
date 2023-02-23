package services

import (
	"encoding/csv"
	"github.com/pkg/errors"
	"os"

	"fiber-base-go/internal/model"
)

type PopulateService interface {
	Populate() error
}

type populateService struct {
	roomService RoomService
}

var _ PopulateService = (*roomService)(nil)

func NewPopulateService(roomService RoomService) PopulateService {
	return &populateService{roomService}
}

func (s *populateService) Populate() error {
	return s.PopulateRoom()
}

func (s *populateService) PopulateRoom() error {
	file, err := os.Open("sample_data/room.csv")
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
