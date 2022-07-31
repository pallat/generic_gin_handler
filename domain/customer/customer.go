package customer

import (
	"example/generics/ginhandlerfunc"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CustomerRequest struct {
	Name string `json:"name"`
}
type CustomerResponse struct {
	Name string `json:"name"`
}

func CustomerHandler(request CustomerRequest) (CustomerResponse, error) {
	fmt.Println("CustomerHandler", request)
	return CustomerResponse{Name: request.Name}, nil
}

func New() gin.HandlerFunc {
	return ginhandlerfunc.NewBindRequestOKResponse(CustomerHandler)
}
