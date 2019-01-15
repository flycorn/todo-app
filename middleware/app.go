package middleware

import (
	"github.com/gin-gonic/gin"
)

func Default() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()
	}
}