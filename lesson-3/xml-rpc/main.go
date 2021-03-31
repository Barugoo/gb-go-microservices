package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://postgres:5432/postgres?sslmode=disable&user=postgres&database=postgres&password=postgres",
	)
	if err != nil {
		log.Fatal(err)
	}
	for db.Ping() != nil {
		log.Println(db.Ping())
	}
	s := &Service{db}

	srv := rpc.NewServer()

	srv.RegisterCodec(json.NewCodec(), "application/json")

	srv.RegisterService(s, "Service")
	http.Handle("/rpc", srv)

	log.Println("Starting server on localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

type Service struct {
	db *sql.DB
}

func (h *Service) GetMovie(r *http.Request, in *struct{ Id int }, out *Movie) error {
	var res Movie
	err := h.db.QueryRow("SELECT id, name, poster, movie_url FROM movies WHERE id = $1", in.Id).
		Scan(&res.ID, &res.Name, &res.Poster, &res.MovieUrl)
	if err != nil {
		log.Println(*in, err)
	}
	*out = res
	return err
}

type Movie struct {
	ID       int
	Name     string
	Poster   string
	MovieUrl string
}
