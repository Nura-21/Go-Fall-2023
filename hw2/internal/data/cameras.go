package data

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"hw2/internal/validator"
	"time"
)

type CameraModel struct {
	DB *sql.DB
}

func (c CameraModel) Insert(camera *Camera) error {
	query := `INSERT INTO cameras (title, year, manufacturer, model, details)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, created_at`
	args := []interface{}{camera.Title, camera.Year, camera.Manufacturer, camera.Model, camera.Details}
	return c.DB.QueryRow(query, args...).Scan(&camera.ID, &camera.CreatedAt)
}

func (c CameraModel) Get(id int64) (*Camera, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, title, year, manufacturer, model, details, version
		FROM cameras
		WHERE id = $1`
	var camera Camera
	err := c.DB.QueryRow(query, id).Scan(
		&camera.ID,
		&camera.CreatedAt,
		&camera.Title,
		&camera.Year,
		&camera.Manufacturer,
		&camera.Model,
		&camera.Details,
		&camera.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &camera, nil
}

func (c CameraModel) Update(camera *Camera) error {
	query := `
		UPDATE cameras
		SET title = $1, year = $2, manufacturer = $3, model = $4, details = $5, version = version + 1
		WHERE id = $6
		RETURNING version`
	args := []interface{}{
		camera.Title,
		camera.Year,
		camera.Manufacturer,
		camera.Model,
		camera.Details,
		camera.ID,
	}
	return c.DB.QueryRow(query, args...).Scan(&camera.Version)
}

func (c CameraModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM cameras
		WHERE id = $1`
	result, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

type Camera struct {
	ID           int64     `json:"id"`           // Unique integer ID for the movie
	CreatedAt    time.Time `json:"-"`            // Timestamp for when the camera is added to our database
	Title        string    `json:"title"`        // Camera title
	Year         int32     `json:"year"`         // Camera release year
	Manufacturer string    `json:"manufacturer"` // Camera manufacturer
	Model        string    `json:"model"`        // Camera model
	Details      string    `json:"details"`      // Camera details
	Version      int32     `json:"version"`      // Camera version
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
