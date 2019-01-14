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
	err := service.Db.Order("id DESC").Find(&users).RecordNotFound()
	if(err){
		c.JSON(200, nil)
	}
	c.JSON(200, users)

	//c.File(moduleViewPath+"index.html")
}