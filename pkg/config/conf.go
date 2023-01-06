package config

import "github.com/spf13/viper"

func Get() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return config
}
