package users

import (
	"github.com/buraksecer/go-clean-architecture/entities"
	"github.com/buraksecer/go-clean-architecture/repositories/mockdatabase"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"testing"
)

func TestSignUp(t *testing.T) {
	type args struct {
		input entities.User
	}

	repository := mockdatabase.NewMockDatabase()
	logger := log.With().Logger()

	ser := LoadService(repository, &logger)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful sign up",
			args: args{input: entities.User{
				FirstName: "burak",
				LastName:  "seçer",
			},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("test SignUp with "+tt.name, func(t *testing.T) {
			err := ser.SignUp(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing SignUp operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	type args struct {
		input entities.User
	}

	repository := mockdatabase.NewMockDatabase()
	logger := log.With().Logger()

	ser := LoadService(repository, &logger)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful sign in",
			args: args{input: entities.User{
				Username: "buraksecer",
				Password: "encrypted-dummy-password",
			},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run("test SignIn with "+tt.name, func(t *testing.T) {
			_, err := ser.SignIn(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing SignIn operation. error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSayHello(t *testing.T) {
	repository := mockdatabase.NewMockDatabase()
	logger := log.With().Logger()

	ser := LoadService(repository, &logger)

	t.Run("test SignIn", func(t *testing.T) {
		c := &gin.Context{}
		c.Set("username", "existing-dummy-user")

		l := entities.User{Username: c.GetString("username")}
		_, err := ser.SayHello(l)
		if err != nil {
			t.Errorf("error while executing SayHello operation. error = %v", err)
		}
	})
}