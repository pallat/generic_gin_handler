package main

import (
	"example/generics/domain/customer"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/customers", customer.New())

	r.Run()
}
