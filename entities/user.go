package entities

import (
	"errors"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (input User) Validate() error {
	switch {
	case input.FirstName == "":
		return errors.New("")
	default:
		return nil
	}
}
