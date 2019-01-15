package api

import (
	"github.com/gin-gonic/gin"
	"todo-app/module/api/controller"
	"todo-app/middleware"
)

/**
 加载模块
 */
func Load(router *gin.RouterGroup, app *gin.Engine){
	//权限白名单
	allowUrl := []string{
		"/api/test",
		"/api/get-users",
	}
	//授权权限
	router.Use(middleware.ApiAuth(allowUrl))

	//获取条目数据
	router.GET("/get-todos", api.GetTodos)

	router.GET("/get-users", api.GetUsers)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 200,
			"message": "auth test",
		})
	})
}