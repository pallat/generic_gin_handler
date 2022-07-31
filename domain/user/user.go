package user

import (
	"example/generics/ginhandlerfunc"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Name string `uri:"name" binding:"required"`
}
type UserResponse struct {
	Name string `json:"name"`
}

func UserHandler(request UserRequest) (UserResponse, error) {
	fmt.Println("UserHandler", request)
	return UserResponse{Name: request.Name}, nil
}

func New() gin.HandlerFunc {
	return ginhandlerfunc.NewParamOKResponse(UserHandler)
}
