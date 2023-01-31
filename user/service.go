package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	RegisterInput(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckExistedUserByEmail(input UniqueEmailInput) (bool, error)
	SaveAvatar(Id int, fileLocation string) (User, error)
}

//// validate interface is implement
//var _ Service = (*service)(nil)

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

func (s *service) Login(input LoginInput) (User, error) {

	existedUser, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return existedUser, err
	}

	errorBcrypt := bcrypt.CompareHashAndPassword([]byte(existedUser.PasswordHash), []byte(input.Password))

	if errorBcrypt != nil {
		return existedUser, err
	}

	return existedUser, nil
}

func (s *service) CheckExistedUserByEmail(input UniqueEmailInput) (bool, error) {
	email := input.Email

	existedUser, err := s.repository.FindByEmail(email)

	if existedUser.Id == 0 {
		return true, err
	}

	return false, nil
}

func (s *service) SaveAvatar(Id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(Id)

	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation
	s.repository.Update(user)

	return user, nil
}
