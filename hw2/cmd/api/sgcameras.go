package main

import (
	"fmt"
	"hw2/internal/data"
	"hw2/internal/validator"
	"net/http"
	"time"
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
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	camera := data.Camera{
		ID:           id,
		CreatedAt:    time.Now(),
		Title:        "Canon Pro",
		Year:         2023,
		Manufacturer: "Canon",
		Model:        "Pro",
		Details:      "Some details",
	}

	err = app.writeJSON(w, http.StatusOK, camera, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
