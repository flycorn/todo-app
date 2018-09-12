package main

import (
	"./helper"
)

func main(){

	helper.Dd("----test----")

	params := make(map[string]string)
	params["name"] = "jack"
	params["id"] = "2"

	str := helper.GenerateToken(params)
	helper.Dd(str)
}
