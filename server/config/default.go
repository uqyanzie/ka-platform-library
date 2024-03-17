package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	NEO4J_URI  string `mapstructure:"NEO4J_URI"`
	NEO4J_USER string `mapstructure:"NEO4J_USER"`
	NEO4J_PASS string `mapstructure:"NEO4J_PASS"`
	// RedisUri string `mapstructure:"REDIS_URL"`
	Port string `mapstructure:"PORT"`

	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`

	Origin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
