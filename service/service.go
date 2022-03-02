package service

import (
	"refactoring/model"
	"refactoring/repository"
)

type User interface {
	Create(user model.CreateUserRequest) (string, error)
	Get(id string) (*model.User, error)
	Update(id string, userUpdate model.UpdateUserRequest) error
	Delete(id string) error
	Search(searchQuery string) model.UserList
}

type Service struct {
	User
}

func NewService(rep *repository.Repository) *Service {
	return &Service{User: NewUserService(rep.User)}
}
