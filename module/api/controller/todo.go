package controller

import (
	"github.com/gin-gonic/gin"
	"todo-app/helper"
	"todo-app/model"
	"todo-app/service"
)

//获取列表
func GetTodos(c *gin.Context){
	//u := helper.GetUserData(c) //获取用户数据

	//status := c.DefaultQuery("status", "1")

}

//添加数据
func PostTodo(c *gin.Context){
	uid := helper.GetUid(c) //获取当前用户数据
	if uid == 0{
		helper.ReturnApi(c, 400, "用户有误")
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	if content == "" {
		helper.ReturnApi(c, 400, "内容不能为空")
		return
	}
	if title == "" {
		//标题为空、截取内容部分内容为标题(默认截取15字)
		title = helper.SubString(content, 0, 15)
	}

	//写入数据
	todo := model.Todo{}
	todo.Uid = uid
	todo.Title = title
	todo.Content = content
	todo.Status = 1
	if err := service.Db.Create(&todo).Error; err != nil {
		//失败
		helper.ReturnApi(c, 400, "提交失败:"+err.Error())
		return
	}
	helper.ReturnApi(c, 200, "提交成功", todo)
}

//更新数据状态
func UpdateTodoStatus(c *gin.Context){

}

//更新数据
func UpdateTodo(c *gin.Context){

}