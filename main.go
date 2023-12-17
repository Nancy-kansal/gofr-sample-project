package main

import (
	"gofr.dev/pkg/gofr"

	ht "Project/http/movie"
	serv "Project/service/movie"
	"Project/store/movie"
)

func main() {
	datastore := movie.NewDBHandler()
	serviceHandler := serv.NewMovieServiceHandler(datastore)
	handler := ht.New(serviceHandler)

	app := gofr.New()
	app.Server.ValidateHeaders = false

	app.GET("/movies/{id}", handler.GetByIDRequest)
	app.GET("/movies", handler.GetAllRequest)

	app.POST("/movies", handler.CreateMovieRequest)
	app.PUT("/movies/{id}", handler.UpdateByIDRequest)

	app.DELETE("/movies/{id}", handler.DeleteByIDRequest)
	app.Start()
}
