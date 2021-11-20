package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"testing"
	"time"
)

func TestValidateToken(t *testing.T) {
	type args struct {
		secret     string
		username   string
		expiration int64
	}
	test := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "successful token validation",
			args:    args{secret: "b1s2", username: "buraksecer", expiration: time.Now().Add(time.Hour * time.Duration(1)).Unix()},
			wantErr: false,
		},
		{
			name:    "unsuccessful token validation due to a bad secret",
			args:    args{secret: "bad-secret", username: "buraksecer", expiration: time.Now().Add(time.Hour * time.Duration(1)).Unix()},
			wantErr: true,
		},
		{
			name:    "unsuccessful token validation due expired token",
			args:    args{secret: "b1s2", username: "buraksecer", expiration: time.Now().Add(time.Hour * time.Duration(-1)).Unix()},
			wantErr: true,
		},
	}

	for _, tt := range test {
		t.Run("test ValidateToken with "+tt.name, func(t *testing.T) {
			token := jwt.New(jwt.GetSigningMethod("HS256"))

			now := time.Now().Local()
			token.Claims = jwt.MapClaims{
				"username": tt.args.username,
				"iat":      now.Unix(),
				"exp":      tt.args.expiration,
			}

			tokenString, _ := token.SignedString([]byte("b1s2"))

			_, err := ValidateToken(tokenString, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("error while executing ValidateToken. error = %v, wantErr %v", err, tt.wantErr)
			} else {
				log.Info().Msg("Pass -" + tt.name)
			}
		})
	}
}
