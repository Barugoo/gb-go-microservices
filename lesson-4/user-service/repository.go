package main

import (
	uuid "github.com/google/uuid"
)

type UserStorage []*User

var UU = UserStorage{
	&User{1, "bob@mail.ru", "Bob", true, "god", "1"},
	&User{2, "alice@mail.ru", "Alice", false, "secret", "2"},
}

func (uu UserStorage) CreateUser(user *User) *User {
	maxID := uu[len(uu)-1].ID
	user.ID = maxID + 1
	user.Token = uuid.New().String()
	UU = append(UU, user)
	return user
}

func (uu UserStorage) GetByEmail(email string) *User {
	for _, u := range uu {
		if u.Email == email {
			return u
		}
	}
	return nil
}

func (uu UserStorage) GetByToken(token string) *User {
	for _, u := range uu {
		if u.Token == token {
			return u
		}
	}
	return nil
}

func (uu UserStorage) GetByID(id int) *User {
	for _, u := range uu {
		if u.ID == id {
			return u
		}
	}
	return nil
}
