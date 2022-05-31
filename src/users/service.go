package users

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterInput) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		fmt.Println("User not available")
		return user, err
	}

	if user.ID != 0 {
		return user, errors.New("User not available")
	}

	inputUser := User{}
	inputUser.Name = input.Name
	inputUser.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	inputUser.Password = string(passwordHash)
	inputUser.PhoneNumber = input.PhoneNumber
	inputUser.Address = input.Address

	newUser, err := s.repository.Save(inputUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
