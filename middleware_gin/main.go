package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//创建日志文件
	f, _ := os.Create("gin.log")
	//修改默认输出器，输出到上面文件
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	//单独使用Logger中间件创建引擎
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		panic("test panic")
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})

	r.Run()
}
