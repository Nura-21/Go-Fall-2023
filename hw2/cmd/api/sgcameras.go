package main

import (
	"fmt"
	"net/http"
)

func (app *application) createSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new smart garden camera")
}

func (app *application) showSGCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of camera %d\n", id)
}
