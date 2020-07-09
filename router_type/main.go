package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	//接收任意类型请求方法
	r.Any("/any", func(context *gin.Context) {
		context.String(200, "any")
	})
	r.Run()
}
