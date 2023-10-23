package models

import (
	"errors"
	"time"

	"github.com/upper/db/v4"
)

type User struct {
	ID        uint64    `db:"id,omitempty"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	Activated bool      `db:"activated"`
}

func (m UsersModel) Table() string {
	return "users"
}

type UsersModel struct {
	db db.Session
}

func (m *UsersModel) Get(id int) (*User, error) {
	var user User
	err := m.db.Collection("users").Find(db.Cond{id: id}).One(&user)
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return nil, ErrNoMoreRows
		}
	}
	return &user, nil
}
