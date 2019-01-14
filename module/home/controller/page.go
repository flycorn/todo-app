package controller

import (
	"github.com/gin-gonic/gin"
	"todo-app/model"
	"todo-app/service"
)

var moduleViewPath string = "module/home/view/"

//首页
func Index(c *gin.Context){
	var users []model.User
	rs := service.Db.Order("id DESC").Find(&users).RecordNotFound()
	c.JSON(200, rs)

	//c.File(moduleViewPath+"index.html")
}