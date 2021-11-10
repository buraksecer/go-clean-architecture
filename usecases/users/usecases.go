package users

import "github.com/buraksecer/go-clean-architecture/entities"

type UseCase interface {
	SignUp(user entities.User) error
	SignIn(user entities.User) (string, error)
	SayHello(user entities.User) (string, error)
}