package usecase

import (
	"hmrbcnto.com/gin-api/entities"
	"hmrbcnto.com/gin-api/repository"
)

type UserUsecase interface {
	CreateUser(user *entities.CreateUserRequest) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(repo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (userUsecase *userUsecase) CreateUser(userData *entities.CreateUserRequest) (*entities.User, error) {
	// Additional business logic as needed

	// Encrypt password with bcrypt here
	user, err := userUsecase.userRepo.CreateUser(userData)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userUsecase *userUsecase) GetAllUsers() ([]entities.User, error) {
	// Business logic as needed!!

	users, err := userUsecase.userRepo.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}