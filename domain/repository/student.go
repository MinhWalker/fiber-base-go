package repository

import "fiber-base-go/domain"

type StudentRepository interface {
	GetAll() ([]domain.Student, error)
	Create(fact *domain.Student) error
}
