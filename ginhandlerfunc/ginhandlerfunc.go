package ginhandlerfunc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBindRequestOKResponse[Request, Response any](handler func(Request) (Response, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := c.BindJSON(&request); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		response, err := handler(request)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
