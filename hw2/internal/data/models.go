package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Cameras interface {
		Insert(camera *Camera) error
		Get(id int64) (*Camera, error)
		Update(camera *Camera) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Cameras: CameraModel{DB: db},
	}
}
