package configuration

import (
	"log"

	"github.com/spf13/viper"
)

func Init(configName string, configType string, configPath string) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error when load configuration", err)
	}
}
