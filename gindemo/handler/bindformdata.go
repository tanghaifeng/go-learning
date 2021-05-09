package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Member struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func BindHandler(c *gin.Context) {
	m := &Member{}
	c.Bind(m)
	fmt.Println(m)
	c.JSON(200, gin.H{
		"code": "ok",
	})
}
