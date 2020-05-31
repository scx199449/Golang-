package main

import (
	"GinDemo/studydemo01/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	/*
		1、创建一个默认的路由引擎
			gin有三种模式，分别为:
				gin.DebugMode 调试模式
				gin.TestMode  测试模式
				gin.ReleaseMode 发布模式
		    默认是debug模式
			注意：设置模式一定要在Default前面
	*/
	gin.SetMode(gin.DebugMode) //设置模式
	engin := gin.Default()

	engin.LoadHTMLFiles("template/users/index.html")

	engin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "posts/index",
		})
	})

	/*engin.GET("users/index", func(c *gin.Context) {
	    c.HTML(http.StatusOK, "users/index.html", gin.H{
	        "title": "users/index",
	    })
	})*/

	//GET请求
	usergroup := engin.Group("/users")
	{
		usergroup.POST("/Login", controller.Login)
	}

	//启动HTTP服务，使用默认服务器
	engin.Run(":8080")
}
