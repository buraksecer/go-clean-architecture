package entities

import (
	"errors"
	"github.com/buraksecer/go-clean-architecture/global_consts"
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
		return errors.New(global_consts.FirstnameCannotBeEmpty)
	case input.LastName == "":
		return errors.New(global_consts.LastnameCannotBeEmpty)
	case input.Username == "":
		return errors.New(global_consts.UsernameCannotBeEmpty)
	case input.Password == "":
		return errors.New(global_consts.PasswordCannotBeEmpty)
	default:
		return nil
	}
}
