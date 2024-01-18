package config

import (
	"github.com/spf13/viper"
)

func MarshallConfig() (err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
