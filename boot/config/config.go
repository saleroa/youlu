package config

import (
	"bbs/config"
	"log"
)

func InitConfig() {
	err := config.MarshallConfig()
	if err != nil {
		log.Printf("read config error: %v", err)
	}
}
