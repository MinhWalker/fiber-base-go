package services

import (
	"fiber-base-go/domain"
	"fiber-base-go/infrastructure/persistence"
)

type StudentService struct {
	Repo persistence.StudentRepository
}

// GetAllStudents return all domain.news
func (s *StudentService) GetAllStudents() ([]domain.Student, error) {
	return s.Repo.GetAll()
}

// AddStudent saves new Student
func (s *StudentService) AddStudent(p domain.Student) error {
	return s.Repo.Create(&p)
}
