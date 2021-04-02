package main

import (
	"context"
	"database/sql"
	"errors"
)

type UserStorage struct {
	db *sql.DB
}

var (
	ErrNotFound = errors.New("not found")
)

func wrapError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
}

func (s *UserStorage) GetByID(ctx context.Context, userID int32) (*User, error) {

	var res User
	if err := s.db.QueryRow("SELECT id, email, name, is_paid, pwd, token FROM users WHERE id = $1", userID).
		Scan(&res.ID, &res.Email, &res.Name, &res.IsPaid, &res.Pwd, &res.Token); err != nil {
		return nil, wrapError(err)
	}
	return &res, nil
}

func (s *UserStorage) GetByEmail(ctx context.Context, email string) (*User, error) {

	var res User
	if err := s.db.QueryRow("SELECT id, email, name, is_paid, pwd, token FROM users WHERE email = $1", email).
		Scan(&res.ID, &res.Email, &res.Name, &res.IsPaid, &res.Pwd, &res.Token); err != nil {
		return nil, wrapError(err)
	}
	return &res, nil
}

func (s *UserStorage) GetByToken(ctx context.Context, token string) (*User, error) {

	var res User
	if err := s.db.QueryRow("SELECT id, email, name, is_paid, pwd, token FROM users WHERE token = $1", token).
		Scan(&res.ID, &res.Email, &res.Name, &res.IsPaid, &res.Pwd, &res.Token); err != nil {
		return nil, wrapError(err)
	}
	return &res, nil
}

func (s *UserStorage) CreateUser(ctx context.Context, user *User) (*User, error) {
	var res User
	if _, err := s.db.Exec("INSERT INTO users (id, email, name, is_paid, pwd, token) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		user.ID, user.Email, user.Name, user.IsPaid, user.Pwd, user.Token); err != nil {
		return nil, err
	}
	return &res, nil
}
