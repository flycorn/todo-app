package controller

import (
	"github.com/gin-gonic/gin"
	"todo-app/helper"
	"todo-app/model"
	"todo-app/service"
	"strconv"
	"fmt"
)

//获取列表
func GetTodos(c *gin.Context){
	uid := helper.GetUid(c) //获取当前用户数据
	if uid == 0{
		helper.ReturnApi(c, 400, "用户有误")
		return
	}
	//查询
	var todos []model.Todo
	db := service.Db
	db = db.Where("uid = ?", uid)

	//分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1")) //页码
	if page < 0 {
		page = 0
	}
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10")) //每页条数
	if pageSize < 0 {
		pageSize = 10
	}
	p_status := c.DefaultQuery("status", "") //状态
	if p_status != ""{
		fmt.Printf("p_status: %s", p_status)
		//基础验证-数据格式转换
		status, err := strconv.Atoi(p_status)
		if status < 0 || status > 2 || err != nil {
			helper.ReturnApi(c, 400, "状态有误~")
			return
		}
		db = db.Where("status = ?", status)
	}
	err := db.Limit(pageSize).Offset((page - 1) * pageSize).Order("field(status, 1, 2, 0)").Order("id desc").Find(&todos).RecordNotFound()
	if err {
		helper.ReturnApi(c, 200, "数据为空", todos)
		return
	}
	helper.ReturnApi(c, 200, "获取成功", todos)
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
	if title == ""{
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
	uid := helper.GetUid(c) //获取当前用户数据
	if uid == 0{
		helper.ReturnApi(c, 400, "用户有误")
		return
	}
	p_id := c.Param("id") //数据ID
	p_status := c.PostForm("status") //状态
	if p_status == ""{
		p_status = "1"
	}
	//基础验证-数据格式转换
	id, err := strconv.Atoi(p_id)
	if id == 0 || err != nil {
		helper.ReturnApi(c, 404, "该数据有误~")
		return
	}
	status, err := strconv.Atoi(p_status)
	if status < 0 || status > 2 || err != nil {
		helper.ReturnApi(c, 400, "状态有误~")
		return
	}
	todo := model.Todo{}
	//查询该数据
	r := service.Db.Where("id = ? and uid = ?", id, uid).First(&todo).RecordNotFound()
	if r {
		//不存在
		helper.ReturnApi(c, 404, "该数据不存在~")
		return
	}
	if todo.Status == status {
		//不存在
		helper.ReturnApi(c, 400, "状态未改动~")
		return
	}
	todo.Status = status;
	//更新状态
	if err := service.Db.Save(&todo).Error; err != nil {
		//失败
		helper.ReturnApi(c, 400, "提交失败:"+err.Error())
		return
	}
	helper.ReturnApi(c, 200, "提交成功", todo)
}

//更新数据
func UpdateTodo(c *gin.Context){
	uid := helper.GetUid(c) //获取当前用户数据
	if uid == 0{
		helper.ReturnApi(c, 400, "用户有误")
		return
	}
	p_id := c.Param("id") //数据ID
	//基础验证-数据格式转换
	id, err := strconv.Atoi(p_id)
	if id == 0 || err != nil {
		helper.ReturnApi(c, 404, "该数据有误~")
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	if content == "" {
		helper.ReturnApi(c, 400, "内容不能为空")
		return
	}
	todo := model.Todo{}
	//查询该数据
	r := service.Db.Where("id = ? and uid = ?", id, uid).First(&todo).RecordNotFound()
	if r {
		//不存在
		helper.ReturnApi(c, 404, "该数据不存在~")
		return
	}
	if title != ""{
		todo.Title = title
	}
	todo.Content = content
	//修改数据
	if err := service.Db.Save(&todo).Error; err != nil {
		//失败
		helper.ReturnApi(c, 400, "修改失败:"+err.Error())
		return
	}
	helper.ReturnApi(c, 200, "修改成功", todo)
}