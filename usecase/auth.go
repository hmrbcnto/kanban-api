package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"hmrbcnto.com/gin-api/config"
	"hmrbcnto.com/gin-api/entities"
	"hmrbcnto.com/gin-api/repository"
)

type AuthUsecase interface {
	Login(email string, password string) (*string, error)
}

type authUsecase struct {
	userRepo repository.UserRepo
}

func NewAuthUsecase(repo repository.UserRepo) AuthUsecase {
	return &authUsecase{
		userRepo: repo,
	}
}

func (au *authUsecase) Login(email string, password string) (*string, error) {
	user, err := au.userRepo.GetUserByEmail(email)

	if err != nil {
		return nil, errors.New("Invalid login")
	}

	// Match passwords if user exists
	byteLoginPassword := []byte(password)
	byteUserPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteUserPassword, byteLoginPassword)

	if err != nil {
		return nil, errors.New("Invalid login")
	}

	token, err := generateJWT(user)

	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	return token, nil
}

func generateJWT(user *entities.User) (*string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(2 * time.Hour)
	claims["user"] = user

	// If it loaded in main, it's not gonna mess things up again
	// Look for ways not to load it again because this is bloody redundant
	config, err := config.LoadConfig()

	tokenString, err := token.SignedString(config.EnvConfig.SecretKey)

	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	return &tokenString, nil
}
