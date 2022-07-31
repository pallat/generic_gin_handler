package product

import (
	"example/generics/ginhandlerfunc"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Name string `json:"name"`
}
type ProductResponse struct {
	Name string `json:"name"`
}

type Product struct {
	Name string
}

type storer interface {
	Save(Product) error
}

type Handler struct {
	store storer
}

func (h *Handler) SaveProduct(request ProductRequest) (ProductResponse, error) {
	fmt.Println("ProductHandler", request)
	return ProductResponse{Name: request.Name}, nil
}

func New(store storer) gin.HandlerFunc {
	h := &Handler{store: store}

	return ginhandlerfunc.NewBindRequestOKResponse(h.SaveProduct)
}
