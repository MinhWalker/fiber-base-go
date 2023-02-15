package repository

import "fiber-base-go/domain"

type FactRepository interface {
	GetAll() ([]domain.Fact, error)
	Create(fact *domain.Fact) error
}