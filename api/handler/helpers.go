package handler

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/buraksecer/go-clean-architecture/entities"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func ErrHandler(err error, c *gin.Context) {
	switch err {
	case entities.ErrUserAlreadyExists:
		c.AbortWithStatusJSON(409, gin.H{"error": err.Error()})
	case entities.ErrInvalidPassword:
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
	case entities.ErrUserDoesNotExist:
		c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
	case entities.ErrInvalidInput:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	default:
		c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
	}
}

func Unmarshal(c *gin.Context, t interface{}, marshalType string) error {
	x, _ := ioutil.ReadAll(c.Request.Body)

	var err error

	switch marshalType {
	case "json":
		err = json.Unmarshal(x, &t)
	case "xml":
		err = xml.Unmarshal(x, &t)
	}

	if err != nil {
		return entities.ErrInvalidInput
	}
	return nil
}

func (input SigninInput) Validate() error {
	if input.Username == "" {
		return errors.New("username must be provided")
	}

	if input.Password == "" {
		return errors.New("password must be provided")
	}
	return nil
}
