package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/upper/db/v4"
)

var (
	ErrNoMoreRows     = errors.New("no record found")
	ErrEmailDuplicate = errors.New("email already exists")
	ErrUserNotActive  = errors.New("you account is not active")
	ErrInvalidLogin   = errors.New("invalid login")
	errDuplicateVotes = errors.New("you have already voted")
)

type Models struct {
	Users UsersModel
	Posts PostsModel
}

func NewModel(db db.Session) Models {
	return Models{
		Users: UsersModel{db: db},
		Posts: PostsModel{db: db},
	}
}

func convertUpperIDToInt(id db.ID) int {
	idType := fmt.Sprintf("%T", id)
	if idType == "int64" {
		return int(id.(uint64))
	}
	return id.(int)
}

func errHasDuplicate(err error, key string) bool {
	str := fmt.Sprintf("ERROR: duplicate key value violates unique constraint %s", key)
	return strings.Contains(err.Error(), str)
}
