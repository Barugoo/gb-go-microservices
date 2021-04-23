package main

import (
	"context"

	pb "movie-service/api"
)

type Service struct{}

func (s *Service) ListMovies(c context.Context, req *pb.ListMoviesRequest) (resp *pb.ListMoviesResponse, err error) {
	resp.Items = make([]*pb.Movie, 0, len(s.movieList()))
	for _, m := range s.movieList() {
		movie := &pb.Movie{
			Id:     m.ID,
			Name:   m.Name,
			Poster: m.Poster,
			Url:    m.MovieUrl,
		}
		resp.Items = append(resp.Items, movie)
	}
	return resp, nil
}

func (s *Service) GetMovie(c context.Context, req *pb.GetMovieRequest) (resp *pb.Movie, err error) {

	m := s.movieList()[req.MovieId]
	resp = &pb.Movie{
		Id:     m.ID,
		Name:   m.Name,
		Poster: m.Poster,
		Url:    m.MovieUrl,
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
