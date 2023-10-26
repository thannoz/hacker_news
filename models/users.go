package models

import (
	"errors"
	"time"

	"github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
)

const passwordCost = 14

type User struct {
	ID        int       `db:"id,omitempty"`
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

func (m UsersModel) FindByEmail(email string) (*User, error) {

	var u User
	err := m.db.Collection(m.Table()).Find(db.Cond{"email": email}).One(&u)
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return nil, ErrNoMoreRows
		}
		return nil, err
	}
	return &u, nil
}

func (m UsersModel) InsertUser(u *User) error {
	newHashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), passwordCost)
	if err != nil {
		return err
	}
	u.Password = string(newHashPassword)
	u.CreatedAt = time.Now()
	col := m.db.Collection(m.Table())
	res, err := col.Insert(u)
	if err != nil {
		switch {
		case errHasDuplicate(err, "users_email_key"):
			return ErrEmailDuplicate
		default:
			return err
		}
	}

	u.ID = convertUpperIDToInt(res.ID())

	return nil
}
