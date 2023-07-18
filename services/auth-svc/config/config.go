package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"application.mode" default:"dev"`
	Maintenance bool   `mapstructure:"application.maintenance" default:"false"`
	ServiceName string `mapstructure:"services.auth-service" default:"auth-svc"`
	ServicePort int    `mapstructure:"services.auth-service.port" default:"28001"`

	RedisHost     string `mapstructure:"cache.configs.redis.host" default:"209.182.237.44"`
	RedisPort     string `mapstructure:"cache.configs.redis.port" default:"6380"`
	RedisPassword string `mapstructure:"cache.configs.redis.password" default:"OppoDev12@"`
}

func New() (Config, error) {
	cfg := Config{}

	viper.SetConfigName("app.config") // name of the config file (without extension)
	viper.AddConfigPath(".")          // path to look for the config file in
	viper.SetConfigType("yaml")       // type of the config file

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return cfg, nil
}
