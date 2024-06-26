package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"

	// "github.com/danielgtaylor/huma/v2/autopatch"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("Huma Really Rocks", version))

	huma.Register(api, huma.Operation{
		OperationID: "healthcheck",
		Method:      http.MethodGet,
		Path:        "/v1/healthcheck",
		Summary:     "check the server health.",
		Description: "check the server health.",
	}, app.healthcheckHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "createMovie",
		Method:        http.MethodPost,
		Path:          "/v1/movies",
		Summary:       "create a new movie",
		Description:   "create a new movie",
		Tags:          []string{"movies"},
		DefaultStatus: http.StatusCreated,
	}, app.createMovieHandler)

	huma.Register(api, huma.Operation{
		OperationID: "getMovie",
		Method:      http.MethodGet,
		Path:        "/v1/movies/{id}",
		Summary:     "get movie by id",
		Description: "get a movie",
		Tags:        []string{"movies"},
	}, app.getMovieHandler)

	huma.Register(api, huma.Operation{
		OperationID: "updateMovie",
		Method:      http.MethodPut,
		Path:        "/v1/movies/{id}",
		Summary:     "update movie by id",
		Description: "update a movie",
		Tags:        []string{"movies"},
	}, app.updateMovieHandler)

	huma.Register(api, huma.Operation{
		OperationID: "DeleteMovie",
		Method:      http.MethodDelete,
		Path:        "/v1/movies/{id}",
		Summary:     "delete movie by id",
		Description: "delete a movie",
		Tags:        []string{"movies"},
	}, app.deleteMovieHandler)

	huma.Register(api, huma.Operation{
		OperationID: "listMovie",
		Method:      http.MethodGet,
		Path:        "/v1/movies",
		Summary:     "list movies",
		Description: "list movie",
		Tags:        []string{"movies"},
	}, app.listMoviesHandler)

	// autopatch.AutoPatch(api)

	return router
}
