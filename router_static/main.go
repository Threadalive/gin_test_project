package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//绑定静态资源路径
	r.Static("/assets", "./assets")
	//绑定静态文件系统
	r.StaticFS("/static", http.Dir("static"))
	//绑定单个静态文件
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.Run()
}
