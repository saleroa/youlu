package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/Login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		log.Printf("启动服务错误：%v", err)
		return
	}
}
