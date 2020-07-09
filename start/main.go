package main

import "github.com/gin-gonic/gin"

func main() {
	//创建gin的实例
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"mseeage": "pong",
		})
	})
	r.Run()
}
