package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/catherineelfrida/mygram-api/internal/model"
	"github.com/catherineelfrida/mygram-api/internal/repository"
	"github.com/catherineelfrida/mygram-api/pkg"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(user model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(&user)
}

func (s *UserService) LoginUser(email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := pkg.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) UpdateUser(userID string, updatedUser model.User) (*model.User, error) {
	return s.userRepo.Update(userID, &updatedUser)
}

func (s *UserService) DeleteUser(userID string) error {
	return s.userRepo.Delete(userID)
}

