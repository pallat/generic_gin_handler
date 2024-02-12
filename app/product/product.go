package product

import (
	"example/generics/app"
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

func (h *Handler) SaveProduct(request ProductRequest) any {
	fmt.Println("ProductHandler", request)
	return app.ResponseBadRequest{
		Message: "badrequest",
		Data:    ProductResponse{Name: request.Name},
	}
}

func New(store storer) gin.HandlerFunc {
	h := &Handler{store: store}

	return app.NewHandler(h.SaveProduct)
}
