package v1

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, Error{msg})
}
