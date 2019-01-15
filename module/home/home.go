package home

import (
	"github.com/gin-gonic/gin"
	"todo-app/module/home/controller"
)

//模块视图路径
var moduleViewPath string = "module/home/view/"

/**
 加载默认模块
 */
func Load(router *gin.RouterGroup, app *gin.Engine){

	//首页
	//router.GET("/", func(c *gin.Context) {
		//c.File(moduleViewPath+"index.html")
	//})


	//首页
	router.GET("/", home.Index)
}
