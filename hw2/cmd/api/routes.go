package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sgcameras", app.createSGCameraHandler)
	router.HandlerFunc(http.MethodGet, "/v1/sgcameras/:id", app.showSGCameraHandler)

	return router
}
