package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Name string `json:"name"`
}
type UserResponse struct {
	Name string `json:"name"`
}

type appContext interface {
	BindJSON(any) error
	AbortWithError(int, error) *gin.Error
	JSON(int, any)
}

func UserHandler(c appContext) {
	var request UserRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println("UserHandler", request)
	c.JSON(http.StatusOK, &UserResponse{Name: request.Name})
}
