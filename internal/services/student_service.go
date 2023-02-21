package services

import (
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/repository"

	"github.com/pkg/errors"
)

type StudentService interface {
	CreateStudent(*model.Student) error
	GetAllStudents() ([]*model.Student, error)
	GetStudent(uint) (*model.Student, error)
	UpdateStudent(*model.Student) error
	DeleteStudent(*model.Student) error
	ImportStudent([]*model.Student) error
}

type studentService struct {
	studentRepo repository.StudentRepository
}

var _ StudentService = (*studentService)(nil)

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &studentService{studentRepo}
}

func (s *studentService) CreateStudent(student *model.Student) error {
	if err := s.studentRepo.Create(student); err != nil {
		return errors.Wrap(err, "studentService.CreateStudent")
	}

	return nil
}

func (s *studentService) ImportStudent(students []*model.Student) error {
	return s.studentRepo.CreateMany(students)
}

func (s *studentService) GetAllStudents() ([]*model.Student, error) {
	students, err := s.studentRepo.GetAllStudent()
	if err != nil {
		return nil, errors.Wrap(err, "studentService.GetAllStudents")
	}

	return students, nil
}

func (s *studentService) GetStudent(id uint) (*model.Student, error) {
	student, err := s.studentRepo.FindOne(id)
	if err != nil {
		return nil, errors.Wrap(err, "studentService.GetStudent")
	}

	return student, nil
}

func (s *studentService) UpdateStudent(student *model.Student) error {
	if err := s.studentRepo.Update(student); err != nil {
		return errors.Wrap(err, "studentService.UpdateStudent")
	}

	return nil
}

func (s *studentService) DeleteStudent(student *model.Student) error {
	if err := s.studentRepo.Delete(student); err != nil {
		return errors.Wrap(err, "studentService.DeleteStudent")
	}

	return nil
}
