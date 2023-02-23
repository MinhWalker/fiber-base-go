package repository

import (
	"fiber-base-go/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateMany(rooms []*model.Room) error
}

type roomRepository struct {
	db *gorm.DB
}

func (r *roomRepository) CreateMany(rooms []*model.Room) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, room := range rooms {
		if err := tx.FirstOrCreate(&room, model.Room{Name: room.Name}).Error; err != nil {
			tx.Rollback()
			return errors.Wrap(err, "RoomRepository.CreateMany")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return errors.Wrap(err, "RoomRepository.CreateMany")
	}
	return nil
}

var _ RoomRepository = (*roomRepository)(nil)

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db}
}
