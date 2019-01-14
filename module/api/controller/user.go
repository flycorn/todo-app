package api

import (
	"github.com/gin-gonic/gin"
	"todo-app/model"
	"todo-app/service"
)

//注册
func Register(c *gin.Context){

}

//登录
func Login(c *gin.Context){

}

//用户
func GetUsers(c *gin.Context){
	var users []model.User
	rs := service.Db.Order("id DESC").Find(&users).RecordNotFound()
	c.JSON(200, rs)
}