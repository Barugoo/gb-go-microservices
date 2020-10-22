package main

import (
	"context"

	pb "github.com/barugoo/gb-go-microservices/lesson-3/movie-service/api"
)

type Service struct{}

func (s *Service) GetMovie(c context.Context, req *pb.GetMovieRequest) (resp *pb.GetMovieResponse, err error) {

	m := s.movieList()[req.MovieId]
	resp = &pb.GetMovieResponse{
		MovieId: m.ID,
		Name:    m.Name,
		Poster:  m.Poster,
		Url:     m.MovieUrl,
	}
	return resp, nil
}

func (s *Service) movieList() []Movie {
	return []Movie{
		Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE"},
		Movie{1, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc"},
		Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY"},
	}
}
