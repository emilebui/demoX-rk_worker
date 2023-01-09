package config

import "github.com/spf13/viper"

func Get(relativePath string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(relativePath)
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return config
}
