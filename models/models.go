package models

import (
	"errors"

	"github.com/upper/db/v4"
)

var (
	ErrNoMoreRows = errors.New("no record found")
)

type Models struct {
	Users UsersModel
}

func NewModel(db db.Session) Models {
	return Models{
		Users: UsersModel{db: db},
	}
}
