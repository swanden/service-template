package repository

import (
	"github.com/swanden/service-template/internal/domain/user/entity"
	"github.com/swanden/service-template/internal/domain/user/repository"
	"github.com/swanden/service-template/pkg/database"
)

type UserRepository struct {
	DB *database.DB
}

func New(db *database.DB) *UserRepository {
	return &UserRepository{DB: db}
}

var _ repository.UserRepositoryInterface = (*UserRepository)(nil)

func (r *UserRepository) All() ([]entity.User, error) {
	var users []entity.User
	result := r.DB.Find(&users)

	return users, result.Error
}

func (r *UserRepository) Save(user *entity.User) error {
	result := r.DB.Create(user)

	return result.Error
}

func (r *UserRepository) HasByEmail(email entity.Email) bool {
	var users []*entity.User
	r.DB.Where("email = ?", email.Value).Find(&users)

	return len(users) > 0
}

func (r *UserRepository) Delete(id string) error {
	var user entity.User
	result := r.DB.Where("id = ?", id).Delete(&user)

	return result.Error
}
