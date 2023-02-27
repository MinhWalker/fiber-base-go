package repository

import (
	"fiber-base-go/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	UpsertUser(userReq *model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (f *userRepository) UpsertUser(userReq *model.User) (*model.User, error) {
	var user *model.User
	err := f.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}).FirstOrCreate(&user, model.User{Email: userReq.Email})

	if err != nil {
		return nil, err.Error
	}

	return user, nil
}
