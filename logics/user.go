package logics

import (
	"api/config"
	"api/models"
	"api/repositories"
	"api/utils"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

func InitUserService(userRepository repositories.IUserRepository) *UserService {
	if utils.IsNil(userRepository) {
		userRepository = repositories.InitUserRepository(config.Connection())
	}

	return &UserService{
		userRepository: userRepository,
	}
}

type IUserService interface {
	Get() (*[]models.User, error)
	// GetByID(id uint) (*models.User, error)
	// Create(payload *models.User) error
	// Update(id uint, payload *models.User) error
	// Delete(id uint) error
}

func (service UserService) Get() (*[]models.User, error) {
	var (
		users = []models.User{}
	)

	getUsers, err := service.userRepository.Get()
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}

	_ = utils.AutoMap(getUsers, users)

	return &users, nil
}
