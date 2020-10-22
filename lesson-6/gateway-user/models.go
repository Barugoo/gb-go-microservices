package main

import "html/template"

type MainPage struct {
	Movies      *[]Movie
	User        User
	MoviesError string
}

var TT struct {
	MovieList *template.Template
	Login     *template.Template
}

type User struct {
	Name   string
	IsPaid bool
}

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Poster   string `json:"poster"`
	MovieUrl string `json:"movie_url"`
	IsPaid   bool   `json:"is_paid"`
}
