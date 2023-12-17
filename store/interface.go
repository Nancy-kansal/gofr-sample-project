package store

import (
	"Project/model"

	"gofr.dev/pkg/gofr"
)

type MovieRequestManager interface {
	DeleteByID(ctx *gofr.Context, id int) error
	CreateMovie(*gofr.Context, *model.MovieModel) (*model.MovieModel, error)
	GetByID(ctx *gofr.Context, id int) (*model.MovieModel, error)
	UpdateByID(ctx *gofr.Context, movieObj *model.MovieModel) (*model.MovieModel, error)
	GetAll(*gofr.Context) (*[]model.MovieModel, error)
}
