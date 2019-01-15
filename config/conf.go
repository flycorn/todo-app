package config

import (
	"io/ioutil"
	"github.com/sirupsen/logrus"
	"encoding/json"
	"flag"
	"fmt"
)

//配置参数
type confParam struct {
	Port string `json:"port"`
	JwtSecret string `json:"jwtSecret"`
	Db DB `json:"db"`
}

//DB配置
type DB struct{
	Mysql string `json:"mysql"`
	TablePrefix string `json:"tablePrefix"`
}

//配置表
var Conf *confParam

//加载配置
func LoadConf(){
	//模式类型
	m := flag.String("m", "pro", "模式类型")
	p := flag.String("p", "", "端口号")
	flag.Parse();

	fmt.Printf("---------------")
	fmt.Printf(*m)
	fmt.Printf("---------------")

	//拼接配置文件
	configFile := "./config/"+*m+".json";

	Conf = &confParam{}
	//读取对应环境配置
	b ,err := ioutil.ReadFile(configFile)
	if err != nil{
		logrus.Fatal("配置文件不存在")
	}
	errC := json.Unmarshal(b, Conf)
	if errC != nil {
		logrus.Fatal("配置文件解析错误", errC)
	}
	//扩展端口号
	if *p != "" {
		Conf.Port = ":"+*p
	}
	logrus.Info("配置参数 ", Conf)
}

//func init() {
//	LoadConf()
//}