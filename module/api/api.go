package api

import (
	"github.com/gin-gonic/gin"
	"../../module/api/controller"
	"github.com/dgrijalva/jwt-go"
	"time"
	"../../helper"
)

/**
 加载模块
 */
func Load(router *gin.RouterGroup, app *gin.Engine){

	router.GET("/get-todos", todos.GetTodos)

	//jwt
	router.GET("/jwt", func(c *gin.Context){

		token := jwt.New(jwt.SigningMethodES256)

		helper.Dd(token)

		claims := make(jwt.MapClaims)

		helper.Dd(claims)

		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims

		hmacSecret := []byte("hello-golang")

		helper.Dd(len(hmacSecret))

		tokenString, err := token.SignedString(hmacSecret)

		if err != nil{
			helper.Dd("----error----")
			helper.Dd(err.Error())
			helper.Dd("-------")
		} else {
			helper.Dd("----ok----")
			helper.Dd(tokenString)
			helper.Dd("-------")
		}
		//c.String(200, tokenString)
	})
}