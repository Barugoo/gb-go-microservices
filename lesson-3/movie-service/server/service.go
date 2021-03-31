package main

import (
	"context"
	"database/sql"
	"log"

	pb "movie-service/server/api"
)

type Service struct {
	*pb.UnimplementedMovieServer

	db *sql.DB
}

func (s *Service) GetMovie(c context.Context, req *pb.GetMovieRequest) (resp *pb.GetMovieResponse, err error) {

	var res Movie
	if err := s.db.QueryRow("SELECT id, name, poster, movie_url FROM movies WHERE id = $1", req.MovieId).
		Scan(&res.ID, &res.Name, &res.Poster, &res.MovieUrl); err != nil {
		log.Println(*req, err)
		return nil, err
	}

	resp = &pb.GetMovieResponse{
		Id:       res.ID,
		Name:     res.Name,
		Poster:   res.Poster,
		MovieUrl: res.MovieUrl,
	}
	return resp, nil
}
