package main

import (
	"github.com/gin-gonic/gin"
)

//ip白名单中间件
func IpAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.2",
			"localhost",
		}
		flag := false
		clientIp := c.ClientIP()
		for _, ip := range ipList {
			if ip == clientIp {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s is not in ipList", clientIp)
			c.Abort()
		}
	}
}
func main() {
	r := gin.Default()
	//自定义IP白名单中间件，r.USE(IpAuthMiddleWare())
	r.Use(IpAuthMiddleWare())

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "access ok")
	})
	r.Run()
}
