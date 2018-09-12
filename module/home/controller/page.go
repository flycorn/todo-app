package home

import "github.com/gin-gonic/gin"

var moduleViewPath string = "module/home/view/"

//首页
func Index(c *gin.Context){
	c.File(moduleViewPath+"index.html")
}