package main

import "errors"

var (
	ErrUserNotFound = errors.New("USER_NOT_FOUND")
	ErrUnauthorized = errors.New("UNAUTHORIZED")
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
