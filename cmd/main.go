package main

import (
	"bbs/boot/config"
	"bbs/boot/db"
	"bbs/boot/logger"
	"bbs/boot/server"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	db.InitDb()
	server.InitServer()
}
