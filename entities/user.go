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
		return errors.New(FirstnameCannotBeEmpty)
	case input.LastName == "":
		return errors.New(LastnameCannotBeEmpty)
	case input.Username == "":
		return errors.New(UsernameCannotBeEmpty)
	case input.Password == "":
		return errors.New(PasswordCannotBeEmpty)
	default:
		return nil
	}
}

const (
	FirstnameCannotBeEmpty string = "firstname cannot be empty"
	LastnameCannotBeEmpty  string = "lastname cannot be empty"
	UsernameCannotBeEmpty  string = "username cannot be empty"
	PasswordCannotBeEmpty  string = "password cannot be empty"
)
