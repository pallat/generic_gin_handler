package main

import (
	"example/generics/app/customer"
	"example/generics/app/product"
	"example/generics/app/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users/:name", user.New())
	r.POST("/customers", customer.New())
	r.POST("/products", product.New(&db{}))

	r.Run()
}

type db struct{}

func (db) Save(product.Product) error {
	return nil
}
