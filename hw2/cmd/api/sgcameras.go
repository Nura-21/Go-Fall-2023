package main

import (
	"fmt"
	"hw2/internal/data"
	"net/http"
	"time"
)

func (app *application) createSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new smart garden camera")
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
