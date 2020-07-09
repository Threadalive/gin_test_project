package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	//非空，数字大于10
	Age     string `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			return
		} else {
			c.String(200, "person:%v", person)
		}
	})
	r.Run()
}
