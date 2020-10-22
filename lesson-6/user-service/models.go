package main

type User struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	IsPaid bool   `json:"is_paid"`
	Pwd    string `json:"-"`
	Token  string `json:"token"`
}