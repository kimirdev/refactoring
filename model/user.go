package model

import (
	"errors"
	"net/mail"
	"time"
)

type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}
	UserList  map[string]User
	UserStore struct {
		Increment int      `json:"increment"`
		List      UserList `json:"list"`
	}
)

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *CreateUserRequest) Validate() error {
	if c.DisplayName == "" {
		return errors.New("empty Display Name")
	}
	if c.Email == "" {
		return errors.New("empty Email")
	}

	if _, err := mail.ParseAddress(c.Email); err != nil {
		return err
	}

	return nil
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (u *UpdateUserRequest) Validate() error {
	if u.DisplayName == "" {
		return errors.New("empty Display Name")
	}
	return nil
}
