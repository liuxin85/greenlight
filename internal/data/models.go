package data

import (
	"database/sql"
	"errors"
)

// Define custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel.
type Models struct {
	Movies MovieModel
	Tokens TokenModel
	Users  UserModel // Add a new Users field.
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Tokens: TokenModel{DB: db},
		Users:  UserModel{DB: db}, // Initialize a new UserModel instance

	}
}
