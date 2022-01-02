package utils

import "github.com/spf13/viper"

type Config struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisUsername string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AutomaticEnv()
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_USERNAME")
	viper.BindEnv("REDIS_PASSWORD")

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return
	// }

	err = viper.Unmarshal(&config)
	return
}
