package users

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	UpdateUser(input UpdateUserInput) (User, error)
	GetUserByID(ID int) (User, error)
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

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that id")
	}

	return user, nil
}

func (s *service) UpdateUser(input UpdateUserInput) (User, error) {
	user, err := s.repository.FindById(input.ID)
	if err != nil {
		return user, err
	}

	user.Email = input.Email

	if input.Name != "" {
		user.Name = input.Name
	}

	if input.Address != "" {
		user.Address = input.Address
	}

	if input.PhoneNumber != "" {
		user.PhoneNumber = input.PhoneNumber
	}

	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}
		user.Password = string(passwordHash)
	}

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
