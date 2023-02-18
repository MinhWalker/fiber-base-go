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
	students := []domain.Student{}
	if err := f.Conn.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (f *StudentRepository) Create(student *domain.Student) error {
	if err := f.Conn.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func (f *StudentRepository) CreateMany(students []*domain.Student) error {
	tx := f.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, student := range students {
		if err := tx.Create(student).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
