package repository

import "refactoring/model"

type User interface {
	Create(user model.CreateUserRequest) (string, error)
	Get(id string) (*model.User, error)
	Update(id string, userUpdate model.UpdateUserRequest) error
	Delete(id string) error
	Search(searchQuery string) model.UserList
}

type Repository struct {
	User
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserJsonDb(),
	}
}
