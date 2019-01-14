package main

import (
	"./helper"
	"todo-app/config"
)

func main(){

	config.LoadConf()

	helper.Dd("----test----")

	helper.Dd(config.Conf.Db.Mysql)

	//params := make(map[string]string)
	//params["name"] = "jack"
	//params["id"] = "2"

	//str := helper.GenerateToken(params)
	//helper.Dd(str)
}
