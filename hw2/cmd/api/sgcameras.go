package main

import (
	"errors"
	"fmt"
	"hw2/internal/data"
	"hw2/internal/validator"
	"net/http"
)

func (app *application) createSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title        string    `json:"title"`
		Year         data.Year `json:"year"`         // Camera release year
		Manufacturer string    `json:"manufacturer"` // Camera manufacturer
		Model        string    `json:"model"`        // Camera model
		Details      string    `json:"details"`      // Camera details
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	camera := &data.Camera{
		Title:        input.Title,
		Year:         int32(input.Year),
		Manufacturer: input.Manufacturer,
		Model:        input.Model,
		Details:      input.Details,
	}
	v := validator.New()

	if data.ValidateCamera(v, camera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Cameras.Insert(camera)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sgcameras/%d", camera.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"camera": camera}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	camera, err := app.models.Cameras.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"camera": camera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	camera, err := app.models.Cameras.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title        string    `json:"title"`
		Year         data.Year `json:"year"`
		Manufacturer string    `json:"manufacturer"`
		Model        string    `json:"model"`
		Details      string    `json:"details"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	camera.Title = input.Title
	camera.Year = int32(input.Year)
	camera.Manufacturer = input.Manufacturer
	camera.Model = input.Model
	camera.Details = input.Details

	v := validator.New()
	if data.ValidateCamera(v, camera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Cameras.Update(camera)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"camera": camera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Cameras.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "camera successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
