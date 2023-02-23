package services

import (
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/repository"
)

type RoomService interface {
	importRooms(rooms []*model.Room) error
}

type roomService struct {
	roomRepo repository.RoomRepository
}

func (s *roomService) Populate() error {
	//TODO implement me
	panic("implement me")
}

func (s *roomService) importRooms(rooms []*model.Room) error {
	return s.roomRepo.CreateMany(rooms)
}

func (s *roomService) PopulateRoom() error {
	//TODO implement me
	panic("implement me")
}

var _ RoomService = (*roomService)(nil)

func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{roomRepo}
}
