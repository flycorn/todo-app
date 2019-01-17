package helper

import (
	"os/exec"
	"os"
	"strings"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"todo-app/config"
	"github.com/gin-gonic/gin"
	"math/rand"
	"crypto/md5"
	"net/http"
	"strconv"
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
	resToken, err := token.SignedString([]byte(config.Conf.JwtSecret))
	if err != nil {
		return ""
	}
	return resToken
}

//api返回格式
type ApiRes struct{
	code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

//接口返回
func ReturnApi(c *gin.Context, code int, msg string, data ... interface{}){
	c.Set("return_code", code)
	c.Set("return_msg", msg)
	c.Set("return_data", data)

	rs := gin.H{
		"status":code,
		"msg":msg}

	l := len(data)
	if l < 1{
		rs["data"] = make([]interface{}, 0)
	} else if l == 1 {
		rs["data"] = data[0]
	} else {
		rs["data"] = data
	}
	c.JSON(http.StatusOK, rs)
	c.Abort()
	return
}

//生成随机字符串
func RandStringBytes(n int) string {
	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//MD5字符串
func Md5String(str string) string{
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

//截取字符串
func SubString(str string, begin, length int) string{
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}

//获取登录的用户数据
func GetUserData(c *gin.Context) jwt.MapClaims{
	data, bool := c.Get("ApiAuth") //获取用户信息
	if !bool {
		return nil
	}
	return data.(jwt.MapClaims)
}

//获取用户ID
func GetUid(c *gin.Context) int {
	u := GetUserData(c) //获取用户数据
	if u == nil{
		return 0
	}
	uidS := u["uid"].(string)
	uid, err := strconv.Atoi(uidS)
	if err != nil{
		return 0
	}
	return uid
}