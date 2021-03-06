package gindemo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learning/gindemo/handler"
	"net/http"
	"time"
)

func GinServer() {
	engine := gin.Default()
	engine.Use(TestMiddleware)
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "test")
	})
	engine.GET("/bind", handler.BindHandler)
	v1 := engine.Group("v1/")
	v1.Use(TestMiddleware1)
	{
		v1.GET("ping", func(context *gin.Context) {
			context.String(http.StatusOK, "成功")
		})
	}
	v2 := engine.Group("v2/")
	v2.Use(TestMiddleware3)
	{
		v2.GET("ping", func(context *gin.Context) {
			context.String(http.StatusOK, "成功1")
		})
	}

	// 自定义http配置
	s := http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    6 * time.Second,
		WriteTimeout:   6 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	//engine.Run(":8080")
}

func TestMiddleware(context *gin.Context) {
	fmt.Println("这是一个简单的中间件1")
}

func TestMiddleware1(context *gin.Context) {
	fmt.Println("这是一个简单的中间件2")
}

func TestMiddleware3(context *gin.Context) {
	fmt.Println("这是一个简单的中间件3")
}
