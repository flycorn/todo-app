package home

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"todo-app/helper"
)

var moduleViewPath string = "module/home/view/"

//首页
func Index(c *gin.Context){
	c.File(moduleViewPath+"index.html")
}

type res struct {
	Status int `json:"status"`
	Msg	string `json:"msg"`
	Data interface{} `json:"data"`
}

//测试
func Test(c *gin.Context){
	var apiUrl = "http://localhost:8081/api/user/login";
	var apiHelper helper.ApiHttp

	//p := map[string]interface{}{"mobile": "17602103435"}
	//s := apiHelper.GetSign(p, "8c8eb15034da41a9ac26e8a68d03eed7", 2)
	//
	//fmt.Println(s)

	err := apiHelper.POST(apiUrl, map[string]interface{}{"nickname": "flycorn", "password": "asdasd"})
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("--------------")
	fmt.Println(apiHelper.Body)
}