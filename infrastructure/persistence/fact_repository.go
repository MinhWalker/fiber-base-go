package persistence

import (
	"fiber-base-go/domain"
	"fiber-base-go/domain/repository"
	"gorm.io/gorm"
)

// FactRepositoryImpl Implements repository.FactRepository
type FactRepositoryImpl struct {
	Conn *gorm.DB
}

// NewNewsRepositoryWithRDB returns initialized NewsRepositoryImpl
func NewNewsRepositoryWithRDB(conn *gorm.DB) repository.FactRepository {
	return &FactRepositoryImpl{Conn: conn}
}

func (f *FactRepositoryImpl) GetAll() ([]domain.Fact, error) {
	facts := []domain.Fact{}
	if err := f.Conn.Find(&facts).Error; err != nil {
		return nil, err
	}
	return facts, nil
}

func (f *FactRepositoryImpl) Create(fact *domain.Fact) error {
	if err := f.Conn.Create(&fact).Error; err != nil {
		return err
	}

	return nil
}
