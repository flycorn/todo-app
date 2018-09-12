package middleware

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"../config"
)

/**
 Api授权验证中间件
 */
func ApiAuth(allowUrl ...[]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//过滤白名单
		if len(allowUrl) > 0 && len(allowUrl[0]) > 0{
			//获取当前接口
			u, err := url.Parse(c.Request.RequestURI)
			if err == nil{
				for _,uri := range allowUrl[0]{
					if uri == u.Path {
						c.Next()
						return
					}
				}
			}
		}
		authType := 0 //权限类型 0 header参数 1 get参数

		//获取token
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == ""{
			//兼容GET请求
			tokenStr = c.DefaultQuery("token", "")
			if tokenStr == ""{
				c.JSON(200, gin.H{
					"status": 401,
					"message": "not token",
				})
				c.Abort()
				return
			}
			authType = 1
		}
		//提示信息
		msgError := "Token is error"

		if authType == 0{
			//针对header参数提取token
			tokenStr = string([]byte(tokenStr)[7:])
		}

		//验证token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			//接口签名
			secretStr := config.CONFIG.JwtSecret
			return []byte(secretStr), nil
		})

		if(err == nil && token != nil){
			//判断验证
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				//成功
				c.Set("AuthData", claims)
				c.Next()
				return
			}
		} else {
			msgError = err.Error()
		}
		//验证失败
		c.JSON(200, gin.H{
			"status": 401,
			"message": msgError,
		})
		c.Abort()
		return
	}
}