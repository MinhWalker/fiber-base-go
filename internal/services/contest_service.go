package services

import (
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/repository"
	"github.com/pkg/errors"
)

type ContestService interface {
	CreateContest(contest *model.Contest) error
}

type contestService struct {
	contestRepo repository.ContestRepository
}

var _ ContestService = (*contestService)(nil)

func NewContestService(contestRepo repository.ContestRepository) ContestService {
	return &contestService{contestRepo}
}

func (c contestService) CreateContest(contest *model.Contest) error {
	if err := c.contestRepo.Create(contest); err != nil {
		return errors.Wrap(err, "contestService.CreateStudent")
	}

	return nil
}
