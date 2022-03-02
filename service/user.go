package service

import (
	"refactoring/model"
	"refactoring/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(user model.CreateUserRequest) (string, error) {
	if err := user.Validate(); err != nil {
		return "", err
	}

	id, err := u.repo.Create(user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *UserService) Get(id string) (*model.User, error) {
	user, err := u.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) Update(id string, userUpdate model.UpdateUserRequest) error {
	if err := userUpdate.Validate(); err != nil {
		return err
	}

	return u.repo.Update(id, userUpdate)
}

func (u *UserService) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *UserService) Search(searchQuery string) model.UserList {
	return u.repo.Search(searchQuery)
}
