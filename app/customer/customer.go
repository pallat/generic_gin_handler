package customer

import (
	"example/generics/app"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CustomerRequest struct {
	Name string `json:"name"`
}
type CustomerResponse struct {
	Name string `json:"name"`
}

func CustomerHandler(request CustomerRequest) any {
	fmt.Println("CustomerHandler", request)
	return app.ResponseInternalServerError{
		Message: "error",
		Data:    CustomerResponse{Name: request.Name},
	}
}

func New() gin.HandlerFunc {
	return app.NewHandler(CustomerHandler)
}
