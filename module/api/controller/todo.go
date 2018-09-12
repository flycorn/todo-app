package api

import "github.com/gin-gonic/gin"

//获取列表
func GetTodos(c *gin.Context){
	c.JSON(200, gin.H{
		"status": 200,
		"message": "hello gin",
	})
}