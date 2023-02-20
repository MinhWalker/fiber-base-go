package repository

import (
	"fiber-base-go/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ContestRepository interface {
	Create(contest *model.Contest) error
}

type contestRepository struct {
	db *gorm.DB
}

func NewContestRepository(db *gorm.DB) ContestRepository {
	return &contestRepository{db}
}

func (c *contestRepository) Create(contest *model.Contest) error {
	if err := c.db.Create(contest).Error; err != nil {
		return errors.Wrap(err, "contestRepository.Create")
	}

	return nil
}
