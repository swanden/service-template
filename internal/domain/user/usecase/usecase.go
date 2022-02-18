package usecase

import (
	"github.com/swanden/service-template/internal/domain/user/entity"
	userErrors "github.com/swanden/service-template/internal/domain/user/errors"
	"github.com/swanden/service-template/internal/domain/user/repository"
	"github.com/swanden/service-template/internal/domain/user/service/password"
)

type CreateDTO struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type UserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func New(repository repository.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		userRepository: repository,
	}
}

func (u *UserUseCase) Create(createDTO CreateDTO) (entity.User, error) {
	email, err := entity.NewEmail(createDTO.Email)
	if err != nil {
		return entity.User{}, err
	}

	if u.userRepository.HasByEmail(*email) {
		return entity.User{}, userErrors.AlreadyExistsUser
	}

	pwd, err := password.Hash(createDTO.Password)
	if err != nil {
		return entity.User{}, err
	}
	name := entity.NewName(createDTO.FirstName, createDTO.LastName)
	user := entity.New(entity.NextID(), email, name, pwd)

	return *user, u.userRepository.Save(user)
}

type DeleteDTO struct {
	ID string
}

func (u *UserUseCase) Delete(deleteDTO DeleteDTO) error {
	return u.userRepository.Delete(deleteDTO.ID)
}
