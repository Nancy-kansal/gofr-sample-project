package service

import (
	"Project/model"

	"gofr.dev/pkg/gofr"
)

type MovieManager interface {
	InsertMovieService(*gofr.Context, *model.MovieModel) (*model.MovieModel, error)
	GetByIDService(*gofr.Context, int) (*model.MovieModel, error)
	DeleteByIDService(ctx *gofr.Context, id int) error
	UpdatedByIDService(*gofr.Context, *model.MovieModel) (*model.MovieModel, error)
	GetAllService(*gofr.Context) (*[]model.MovieModel, error)
}
