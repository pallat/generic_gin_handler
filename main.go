package main

import (
	"example/generics/domain/customer"
	"example/generics/domain/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users/:name", user.New())
	r.POST("/customers", customer.New())

	r.Run()
}
