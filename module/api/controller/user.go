package controller

import (
	"github.com/gin-gonic/gin"
	"todo-app/helper"
	"todo-app/model"
	"strconv"
)

//注册
func Register(c *gin.Context){
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	if nickname == ""{
		helper.ReturnApi(c, 400, "昵称不能为空")
		return
	}
	//验证昵称合法长度
	if len(nickname) < 3 || len(nickname) > 10{
		helper.ReturnApi(c, 400, "昵称合法长度(3~10)")
		return
	}
	if password == ""{
		helper.ReturnApi(c, 400, "密码不能为空")
		return
	}
	//验证密码合法长度
	if len(password) < 6 || len(password) > 10{
		helper.ReturnApi(c, 400, "密码合法长度(6~10)")
		return
	}
	//创建用户
	user := model.User{}
	user.Nickname = nickname
	user.CreatePassword(password) //生成密码

	uid, rs := user.CreateUser()
	if rs != nil{
		helper.ReturnApi(c, 500, rs.Error())
		return
	}
	user.Id = uid
	helper.ReturnApi(c, 200, "注册成功", tokenData(user))
}

//登录
func Login(c *gin.Context){
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")

	if nickname == ""{
		helper.ReturnApi(c, 400, "昵称不能为空")
		return
	}
	if password == ""{
		helper.ReturnApi(c, 400, "密码不能为空")
		return
	}

	//查询该用户是否存在
	user := model.User{}
	rs := user.FindUser(nickname)
	if !rs {
		helper.ReturnApi(c, 404, "该用户不存在")
		return
	}
	//验证密码
	rs = user.CheckPassword(password)
	if !rs {
		helper.ReturnApi(c, 400, "密码有误")
		return
	}
	helper.ReturnApi(c, 200, "登录成功", tokenData(user))
}

//格式化Token数据
func tokenData(user model.User) map[string]string{
	//生成jwt Token
	params := make(map[string]string)
	params["uid"] = strconv.Itoa(user.Id)
	params["nickname"] = user.Nickname
	token := helper.GenerateToken(params)

	res := make(map[string]string)
	res["token"] = token
	return res
}