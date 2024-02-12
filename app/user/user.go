package user

import (
	"example/generics/app"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Name string `uri:"name"`
}
type UserResponse struct {
	Name string `json:"name"`
}

func UserHandler(request UserRequest) any {
	fmt.Println("UserHandler", request)
	return app.ResponseOK{
		Message: "ok",
		Data:    UserResponse{Name: request.Name},
	}
}

func New() gin.HandlerFunc {
	return app.NewHandler(UserHandler)
}
