package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
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
	// Check if user with email already exists
	user, _ := userUsecase.userRepo.GetUserByEmail(userData.Email)

	if user != nil {
		return nil, errors.New("Email already in use")
	}

	// Additional business logic as needed

	// Encrypt password with bcrypt here
	bytePassword := []byte(userData.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	user, err = userUsecase.userRepo.CreateUser(&entities.CreateUserRequest{
		Email:    userData.Email,
		Password: string(hashedPassword),
		Username: userData.Username,
	})

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
