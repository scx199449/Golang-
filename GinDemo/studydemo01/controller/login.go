package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("username is :%s\n password is :%s\n", username, password)
	if username != "" && password != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"status":  true,
			"message": "login success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"status":  false,
			"message": "login fail",
		})
	}

}
