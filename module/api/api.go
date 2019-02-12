package api

import (
	"github.com/gin-gonic/gin"
	"todo-app/middleware"
	"todo-app/helper"
	apicon "todo-app/module/api/controller"
	"fmt"
)

//状态码
const CODE_SUCCESS = 200
const CODE_NOT_FIND = 404
const CODE_NO_AUTH = 401
const CODE_ERROR_BAD = 400
const CODE_ERROR_BAN = 403
const CODE_ERROR_SERVER = 500

/**
 加载模块
 */
func Load(router *gin.RouterGroup, app *gin.Engine){
	//权限白名单
	allowUrl := []string{
		"/api/test",
		"/api/user",
		"/api/user/login",
	}
	//授权权限
	router.Use(middleware.ApiAuth(allowUrl))

	//注册
	router.POST("/user", apicon.Register)

	//登录
	router.POST("/user/login", apicon.Login)



	//获取列表数据
	router.GET("/todos", apicon.GetTodos)

	//添加数据
	router.POST("/todo", apicon.PostTodo)

	//更新状态
	router.PUT("/todo-status/:id", apicon.UpdateTodoStatus)

	//修改数据内容
	router.PUT("/todo/:id", apicon.UpdateTodo)


	//测试用户
	router.GET("/user/test", func(c *gin.Context) {
		tmp, _ := c.Get("ApiAuth")
		helper.ReturnApi(c, CODE_SUCCESS, "测试", tmp)
	})

	//测试路由
	router.Any("/test/:id", func(c *gin.Context) {
		id := c.Param("id") //数据ID
		p_id := c.DefaultQuery("id", "1") //状态
		fmt.Printf("id: %s", id)
		fmt.Printf("p_id: %s", p_id)
		return
		//c.JSON(200, gin.H{
		//	"status": 200,
		//	"message": "auth test",
		//})
	})
}