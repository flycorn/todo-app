package helper

import (
	"os/exec"
	"os"
	"strings"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"../config"
)

//获取当前路径
func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

//打印数据
func Dd(d interface{}){
	fmt.Println()
	fmt.Printf("%v", d)
	fmt.Println()
}

//生成TOKEN
func GenerateToken(params map[string]string, extend ...map[string]string) string{
	token := jwt.New(jwt.SigningMethodHS256)
	//数据参数
	claims := make(jwt.MapClaims)
	//默认参数
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	//基础参数
	if len(claims) > 0{
		for k, v := range params {
			claims[k] = v
		}
	}
	//扩展参数
	claims["extend"] = extend
	token.Claims = claims
	//token
	resToken, err := token.SignedString([]byte(config.CONFIG.JwtSecret))
	if err != nil {
		return ""
	}
	return resToken
}