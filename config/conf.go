package config

import (
	"io/ioutil"
	"github.com/sirupsen/logrus"
	"encoding/json"
	"flag"
)

//配置参数
type ConfParam struct {
	Port string `json:"port"`
	JwtSecret string `json:"jwtSecret"`
}

//配置表
var Conf *ConfParam

func LoadConf(){
	//模式类型
	m := flag.String("m", "pro", "模式类型")
	flag.Parse();


	//拼接配置文件
	configFile := "./config/"+*m+".json";

	Conf = &ConfParam{}
	//读取对应环境配置
	b ,err := ioutil.ReadFile(configFile)
	if err != nil{
		logrus.Fatal("配置文件不存在")
	}
	errC := json.Unmarshal(b, Conf)
	if errC != nil {
		logrus.Fatal("配置文件解析错误", errC)
	}
	logrus.Info("配置参数 ", Conf)
}