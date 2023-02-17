package persistence

import (
	"fiber-base-go/domain"
	"fiber-base-go/domain/repository"
	"gorm.io/gorm"
)

// StudentRepositoryImpl Implements repository.StudentRepository
type StudentRepositoryImpl struct {
	Conn *gorm.DB
}

// NewStudentRepositoryWithRDB returns initialized NewsRepositoryImpl
func NewStudentRepositoryWithRDB(conn *gorm.DB) repository.StudentRepository {
	return &StudentRepositoryImpl{Conn: conn}
}

func (f *StudentRepositoryImpl) GetAll() ([]domain.Student, error) {
	facts := []domain.Student{}
	if err := f.Conn.Find(&facts).Error; err != nil {
		return nil, err
	}
	return facts, nil
}

func (f *StudentRepositoryImpl) Create(fact *domain.Student) error {
	if err := f.Conn.Create(&fact).Error; err != nil {
		return err
	}

	return nil
}
