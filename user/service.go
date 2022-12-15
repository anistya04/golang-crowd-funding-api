package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	RegisterInput(input RegisterInput) (User, error)
}

// // validate interface is implement
var _ Service = (*service)(nil)

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterInput(input RegisterInput) (User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user := User{}

	if err != nil {
		return user, nil
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.PasswordHash = string(password)
	user.Role = "user"
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, nil
	}

	return newUser, nil
}
