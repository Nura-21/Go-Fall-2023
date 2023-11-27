package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sgcameras", app.requirePermission("cameras:write", app.createSGCameraHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sgcameras", app.requirePermission("cameras:read", app.listCamerasHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sgcameras/:id", app.requirePermission("cameras:read", app.showSGCameraHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/sgcameras/:id", app.requirePermission("cameras:write", app.updateSGCameraHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/sgcameras/:id", app.requirePermission("cameras:write", app.deleteSGCameraHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
