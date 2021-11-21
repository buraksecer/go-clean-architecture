package handler

import (
	"github.com/buraksecer/go-clean-architecture/api/middleware"
	"github.com/buraksecer/go-clean-architecture/entities"
	"github.com/buraksecer/go-clean-architecture/usecases/users"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	Usecase users.UseCase
}

func NewGinHandler(usecase users.UseCase, jwtSecret string) *gin.Engine {
	h := &GinHandler{
		Usecase: usecase,
	}

	r := gin.Default()
	//r.POST("/users/signin", h.signIn)
	//r.POST("/users/signup", h.signUp)
	r.GET("/hello", middleware.TokenAuthMiddleware(jwtSecret), h.sayHello)
	return r
}

func (h *GinHandler) sayHello(c *gin.Context) {
	var l entities.User
	l.Username = c.GetString("username")

	message, err := h.Usecase.SayHello(l)
	if err != nil {
		ErrHandler(err, c)
		return
	}

	res := HelloOutput{Message: message}

	c.JSON(200, res)
}
