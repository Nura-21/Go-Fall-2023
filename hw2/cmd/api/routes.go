package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sgcameras", app.createSGCameraHandler)
	router.HandlerFunc(http.MethodGet, "/v1/sgcameras/:id", app.showSGCameraHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/sgcameras/:id", app.updateSGCameraHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/sgcameras/:id", app.deleteSGCameraHandler)
	return router
}
