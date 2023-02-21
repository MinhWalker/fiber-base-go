package services

import (
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/repository"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ContestService interface {
	CreateContest(contestName string, grades, schools []string, isAllSchool bool) (*model.Contest, error)
}

type contestService struct {
	contestRepo repository.ContestRepository
	studentRepo repository.StudentRepository
}

var _ ContestService = (*contestService)(nil)

func NewContestService(contestRepo repository.ContestRepository, studentRepo repository.StudentRepository) ContestService {
	return &contestService{
		contestRepo: contestRepo,
		studentRepo: studentRepo,
	}
}

func (c *contestService) CreateContest(contestName string, grades, classes []string, isAllSchool bool) (*model.Contest, error) {
	var students []*model.Student
	var err error
	if isAllSchool {
		students, err = c.studentRepo.GetAllStudent()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrap(err, "contestService.CreateContest")
			}
			return nil, errors.Wrap(err, "contestService.CreateContest")
		}
	} else if len(classes) > 0 {
		students, err = c.studentRepo.FindStudentByClass(classes)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrap(err, "contestService.CreateContest")
			}
			return nil, errors.Wrap(err, "contestService.CreateContest")
		}
	} else if len(grades) > 0 {
		students, err = c.studentRepo.FindStudentByGrade(grades)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrap(err, "contestService.CreateContest")
			}
			return nil, errors.Wrap(err, "contestService.CreateContest")
		}
	} else {
		return nil, errors.New("Don't have any student to create contest")
	}

	contest := &model.Contest{
		Name:     contestName,
		Students: students,
	}

	if err = c.contestRepo.Create(contest); err != nil {
		return nil, errors.Wrap(err, "contestService.CreateContest")
	}

	return contest, nil
}
