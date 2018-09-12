package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Default() gin.HandlerFunc{
	fmt.Println("--------- middleware ----------")
	return func(c *gin.Context){
		fmt.Println("hello : ----" )
		c.Next()
	}
}