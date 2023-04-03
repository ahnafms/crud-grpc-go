package services

import (
	"context"

	"github.com/ahnafms/crud-grpc-go/pkg/database"
	moviePb "github.com/ahnafms/crud-grpc-go/pkg/protobuf"
)

type MovieService struct {
  moviePb.UnimplementedMovieServiceServer
}

func (p *MovieService) CreateMovie(ctx context.Context, req *moviePb.CreateMovieRequest) (*moviePb.CreateMovieResponse, error){
  movie := req.Movie
  database.DB.Create(&movie) 
  return &moviePb.CreateMovieResponse{Status: 1}, nil
}

func (p *MovieService) DeleteMovie(ctx context.Context, req *moviePb.DeleteMovieRequest) (*moviePb.DeleteMovieResponse, error){
  id := req.Id
  database.DB.Delete(&moviePb.Movie{}, &id)
  return &moviePb.DeleteMovieResponse{Status: 1}, nil
}

func (p *MovieService) ListMovies(context.Context, *moviePb.ListMoviesRequest) (*moviePb.ListMoviesResponse, error){
  movies := []*moviePb.Movie{}
  database.DB.Find(&movies)
  return &moviePb.ListMoviesResponse{ MovieList: movies }, nil  
}

func (p *MovieService) UpdateMovie(ctx context.Context, req *moviePb.UpdateMovieRequest) (*moviePb.UpdateMovieResponse, error){
  newMovie := req.Movie
  database.DB.Save(&newMovie)
  return &moviePb.UpdateMovieResponse{Status : 1}, nil
}
