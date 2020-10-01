package main

type MainPage struct {
	Movies *[]Movie
	User   User
	PayURL string
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsPaid bool   `json:"is_paid"`
}

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Poster   string `json:"poster"`
	MovieUrl string `json:"movie_url"`
	IsPaid   bool   `json:"is_paid"`
}
