package data

import (
	"encoding/json"
	"fmt"
	"hw2/internal/validator"
	"time"
)

type Camera struct {
	ID           int64     `json:"id"`           // Unique integer ID for the movie
	CreatedAt    time.Time `json:"-"`            // Timestamp for when the camera is added to our database
	Title        string    `json:"title"`        // Camera title
	Year         int32     `json:"year"`         // Camera release year
	Manufacturer string    `json:"manufacturer"` // Camera manufacturer
	Model        string    `json:"model"`        // Camera model
	Details      string    `json:"details"`      // Camera details
}

func (c Camera) MarshalJSON() ([]byte, error) {
	var year string
	if c.Year != 0 {
		year = fmt.Sprintf("%d year", c.Year)
	}

	aux := struct {
		ID           int64  `json:"id"`
		Title        string `json:"title"`
		Year         string `json:"year"`
		Manufacturer string `json:"manufacturer"`
		Model        string `json:"model"`
		Details      string `json:"details"`
	}{
		ID:           c.ID,
		Title:        c.Title,
		Year:         year,
		Manufacturer: c.Manufacturer,
		Model:        c.Model,
		Details:      c.Details,
	}

	return json.Marshal(aux)
}

func ValidateCamera(v *validator.Validator, camera *Camera) {
	v.Check(camera.Title != "", "title", "must be provided")
	v.Check(len(camera.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(camera.Year != 0, "year", "must be provided")
	v.Check(camera.Year >= 1888, "year", "must be greater than 1888")
	v.Check(camera.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(camera.Manufacturer != "", "manufacturer", "must be provided")
	v.Check(len(camera.Manufacturer) <= 500, "manufacturer", "must not be more than 500 bytes long")
	v.Check(camera.Model != "", "model", "must be provided")
	v.Check(len(camera.Model) <= 500, "model", "must not be more than 500 bytes long")
}
