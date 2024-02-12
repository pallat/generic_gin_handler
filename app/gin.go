package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler[Request, Response any](handler func(Request) Response) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		response := handler(request)

		c.JSON(ResponseStatus(response))
	}
}
