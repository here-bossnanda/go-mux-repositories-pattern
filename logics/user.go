package logics

import (
	"api/app/config"
	"api/entitites"
	"api/models"
	"api/repositories"
	"api/utils"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

func InitUserService(userRepository repositories.IUserRepository) *UserService {
	if utils.IsNil(userRepository) {
		userRepository = repositories.InitUserRepository(config.DB)
	}

	return &UserService{
		userRepository: userRepository,
	}
}

type IUserService interface {
	Get() (*[]models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(payload *models.User) error
	Update(id uint, payload *models.User) error
	Delete(id uint) error
}

func (service UserService) Get() (*[]models.User, error) {
	var (
		users = &[]models.User{}
	)

	getUsers, err := service.userRepository.Get()
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}

	_ = utils.AutoMap(getUsers, users)

	return users, nil
}

func (service UserService) GetByID(id uint) (*models.User, error) {
	var (
		user = &models.User{}
	)

	getUser, err := service.userRepository.GetByID(id)
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}
	_ = utils.AutoMap(getUser, user)

	return user, nil
}

func (service UserService) Create(payload *models.User) error {
	var (
		userEntities = &entitites.User{}
		_            = utils.AutoMap(payload, userEntities)
	)

	err := service.userRepository.Create(userEntities)
	if err != nil {
		go utils.PrintLog(err)
		return err
	}

	return nil
}

func (service UserService) Update(id uint, payload *models.User) error {
	var (
		userEntities = &entitites.User{}
		_            = utils.AutoMap(payload, userEntities)
	)

	_, err := service.GetByID(id)
	if err != nil {
		go utils.PrintLog(err)
		return err
	}

	err = service.userRepository.Update(id, userEntities)
	if err != nil {
		go utils.PrintLog(err)
		return err
	}

	return nil
}

func (service UserService) Delete(id uint) error {
	_, err := service.GetByID(id)
	if err != nil {
		go utils.PrintLog(err)
		return err
	}

	err = service.userRepository.Delete(id)
	if err != nil {
		go utils.PrintLog(err)
		return err
	}

	return nil
}
