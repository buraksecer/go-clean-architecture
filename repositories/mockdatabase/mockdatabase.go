package mockdatabase

import (
	"fmt"
	"github.com/buraksecer/go-clean-architecture/entities"
	"golang.org/x/crypto/bcrypt"
)

type MockDatabase struct{}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{}
}

func (r *MockDatabase) AddUser(details entities.User) error {
	users := []string{"existing-dummy-user"}
	for _, a := range users {
		if a == details.Username {
			return fmt.Errorf("user already exists")
		}
	}
	return nil
}

func (r *MockDatabase) GetUser(username string) (details entities.User, err error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("encrypted-dummy-password"), bcrypt.MinCost)
	existingUser := entities.User{
		FirstName: "burak",
		LastName:  "seçer",
		Password:  string(hash),
		Username:  "buraksecer",
	}

	if username == existingUser.Username {
		return existingUser, nil
	}

	return details, nil
}
