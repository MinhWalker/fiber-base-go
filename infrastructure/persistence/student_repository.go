package persistence

import (
	"fiber-base-go/domain"
	"gorm.io/gorm"
)

// StudentRepository Implements repository.StudentRepository
type StudentRepository struct {
	Conn *gorm.DB
}

func (f *StudentRepository) GetAll() ([]domain.Student, error) {
	facts := []domain.Student{}
	if err := f.Conn.Find(&facts).Error; err != nil {
		return nil, err
	}
	return facts, nil
}

func (f *StudentRepository) Create(fact *domain.Student) error {
	if err := f.Conn.Create(&fact).Error; err != nil {
		return err
	}

	return nil
}
