package repository

import "github.com/swanden/service-template/internal/domain/user/entity"

type UserRepositoryInterface interface {
	All() ([]entity.User, error)
	Save(user *entity.User) error
	HasByEmail(email entity.Email) bool
	Delete(id string) error
}
