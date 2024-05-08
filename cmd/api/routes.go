package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
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

	return router
}
