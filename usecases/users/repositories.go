package users

import "github.com/buraksecer/go-clean-architecture/entities"

type Repository interface {
	GetUser(key string) (entities.User, error)
	AddUser(item entities.User) (err error)
}
