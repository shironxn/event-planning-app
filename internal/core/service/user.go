package service

import (
	"errors"
	"ghepa/internal/core/domain"
	"ghepa/internal/core/port"
	"ghepa/internal/util"
)

type UserService struct {
	repository port.UserRepository
	util       util.Bcrypt
}

func NewUserService(repository port.UserRepository) port.UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Create(req domain.UserRequest) (*domain.User, error) {
	hashedPassword, err := u.util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPassword)

	data, err := u.repository.Create(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Login(req domain.LoginRequest) (*domain.User, error) {
	data, err := u.repository.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := u.util.ComparePassword(req.Password, []byte(data.Password)); err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) GetAll() ([]domain.User, error) {
	data, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) GetByID(req domain.UserRequest) (*domain.User, error) {
	data, err := u.repository.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Update(req domain.UserRequest, claims domain.Claims) (*domain.User, error) {
	user, err := u.repository.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	if user.ID != claims.ID {
		return nil, errors.New("user does not have permission to perform this action")
	}

	hashedPassword, err := u.util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPassword)

	data, err := u.repository.Update(user, req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Delete(req domain.UserRequest, claims domain.Claims) error {
	user, err := u.repository.GetByID(req.ID)
	if err != nil {
		return err
	}

	if user.ID != claims.ID {
		return errors.New("user does not have permission to perform this action")
	}

	err = u.repository.Delete(user)
	if err != nil {
		return err
	}

	return nil

}
