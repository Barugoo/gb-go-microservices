package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Addr = ":8084"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", movieListHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", movieGetHandler).Methods("GET")

	log.Printf("Starting on port %s", Addr)
	log.Fatal(http.ListenAndServe(Addr, r))
}

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Poster   string `json:"poster"`
	MovieUrl string `json:"movie_url"`
	IsPaid   bool   `json:"is_paid"`
}

func movieGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	movieID := vars["id"]

	mm := []Movie{
		Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
		Movie{1, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
		Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	}

	var resultMovie Movie
	for _, movie := range mm {
		if strconv.Itoa(movie.ID) == movieID {
			resultMovie = movie
		}
	}
	if resultMovie.ID == 0 {
		log.Printf("Render response error: %v", errors.New("movie not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(&resultMovie)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func movieListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	mm := []Movie{
		Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
		Movie{1, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
		Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	}
	err := json.NewEncoder(w).Encode(mm)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
