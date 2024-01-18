package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitServer() {
	r := gin.Default()
	err := r.Run(":8080")
	if err != nil {
		log.Printf("gin run error: %v", err)
	}
}
