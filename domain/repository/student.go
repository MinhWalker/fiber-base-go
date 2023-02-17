package repository

import "fiber-base-go/domain"

type FactRepository interface {
	GetAll() ([]domain.Student, error)
	Create(fact *domain.Student) error
}
