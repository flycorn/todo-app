package main

import (
	"github.com/gin-gonic/gin"
	"todo-app/module/home"
	"todo-app/config"
	"todo-app/service"
)

//初始化
func init(){
	//加载配置
	config.LoadConf()
	//加载DB
	service.LoadDb()
}

func main(){
	app := gin.Default()

	//加载静态资源
	app.Static("/assets", "./static")

	//中间件
	//app.Use(middlewares.Default())

	//加载默认模块
	home.Load(app.Group(""), app)

	//加载其它模块~
	//api.Load(app.Group("/api"), app)

	//404
	app.NoRoute(func(c *gin.Context){
		c.String(404, "页面不见喽~~")
	})

	//开启
	app.Run(config.Conf.Port)
}
