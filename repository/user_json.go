package repository

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"refactoring/model"
	"strconv"
	"strings"
	"time"
)

type UserJsonDb struct {
	store string
}

var (
	ErrUserNotFound = errors.New("user_not_found")
)

func NewUserJsonDb() *UserJsonDb {
	return &UserJsonDb{store: `users.json`}
}

func (u *UserJsonDb) Create(user model.CreateUserRequest) (string, error) {
	f, _ := ioutil.ReadFile(u.store)
	s := model.UserStore{}
	_ = json.Unmarshal(f, &s)

	s.Increment++
	us := model.User{
		CreatedAt:   time.Now(),
		DisplayName: user.DisplayName,
		Email:       user.Email,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = us

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(u.store, b, fs.ModePerm)

	return id, nil
}

func (u *UserJsonDb) Get(id string) (*model.User, error) {
	f, _ := ioutil.ReadFile(u.store)
	s := model.UserStore{}
	_ = json.Unmarshal(f, &s)

	user, ok := s.List[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

func (u *UserJsonDb) Update(id string, userUpdate model.UpdateUserRequest) error {
	f, _ := ioutil.ReadFile(u.store)
	s := model.UserStore{}
	_ = json.Unmarshal(f, &s)

	if _, ok := s.List[id]; !ok {
		return ErrUserNotFound
	}

	us := s.List[id]
	us.DisplayName = userUpdate.DisplayName
	s.List[id] = us

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(u.store, b, fs.ModePerm)
	return nil
}

func (u *UserJsonDb) Delete(id string) error {
	f, _ := ioutil.ReadFile(u.store)
	s := model.UserStore{}
	_ = json.Unmarshal(f, &s)

	if _, ok := s.List[id]; !ok {
		return ErrUserNotFound
	}

	delete(s.List, id)

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(u.store, b, fs.ModePerm)

	return nil
}

func (u *UserJsonDb) Search(searchQuery string) model.UserList {
	f, _ := ioutil.ReadFile(u.store)
	s := model.UserStore{}
	_ = json.Unmarshal(f, &s)

	for id, user := range s.List {
		if !strings.Contains(user.DisplayName, searchQuery) && !strings.Contains(user.Email, searchQuery) {
			delete(s.List, id)
		}
	}

	return s.List
}
