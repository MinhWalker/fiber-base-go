package repository

import (
	"gorm.io/gorm"

	"fiber-base-go/internal/model"

	"github.com/pkg/errors"
)

type StudentRepository interface {
	GetAllStudent() ([]*model.Student, error)
	Create(student *model.Student) error
	CreateMany(students []*model.Student) error
	Delete(s *model.Student) error
	FindOne(uint) (*model.Student, error)
	Update(*model.Student) error
	FindStudentByClass([]string) ([]*model.Student, error)
	FindStudentByGrade([]string) ([]*model.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

var _ StudentRepository = (*studentRepository)(nil)

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (f *studentRepository) FindStudentByClass(class []string) ([]*model.Student, error) {
	var students []*model.Student
	result := f.db.Where("class IN (?)", class).Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

func (f *studentRepository) FindStudentByGrade(grades []string) ([]*model.Student, error) {
	var students []*model.Student
	result := f.db.Where("SUBSTRING(class, 1, 2) IN (?)", grades).Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}

	return students, nil
}

func (f *studentRepository) GetAllStudent() ([]*model.Student, error) {
	var students []*model.Student

	if err := f.db.Find(&students).Error; err != nil {
		return nil, errors.Wrap(err, "studentRepository.FindAll")
	}

	return students, nil
}

func (f *studentRepository) FindOne(id uint) (*model.Student, error) {
	var student model.Student

	if err := f.db.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "studentRepository.FindOne")
	}

	return &student, nil
}

func (f *studentRepository) Update(s *model.Student) error {
	if err := f.db.Save(s).Error; err != nil {
		return errors.Wrap(err, "studentRepository.Update")
	}

	return nil
}

func (f *studentRepository) Create(student *model.Student) error {
	if err := f.db.Create(student).Error; err != nil {
		return errors.Wrap(err, "studentRepository.Create")
	}

	return nil
}

func (f *studentRepository) Delete(s *model.Student) error {
	if err := f.db.Delete(s).Error; err != nil {
		return errors.Wrap(err, "studentRepository.Delete")
	}

	return nil
}

func (f *studentRepository) CreateMany(students []*model.Student) error {
	tx := f.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, student := range students {
		if err := tx.Create(&student).Error; err != nil {
			tx.Rollback()
			return errors.Wrap(err, "studentRepository.CreateMany")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return errors.Wrap(err, "studentRepository.CreateMany")
	}

	return nil
}
