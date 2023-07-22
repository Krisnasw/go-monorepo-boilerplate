package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"application.mode" default:"dev"`
	Maintenance bool   `mapstructure:"application.maintenance" default:"false"`
	ServiceName string `mapstructure:"services.auth-service" default:"auth-svc"`
	ServicePort int    `mapstructure:"services.auth-service.port" default:"28001"`

	DBHost         string `mapstructure:"db.host" default:"localhost"`
	DBPort         int    `mapstructure:"db.port" default:"5423"`
	DBUserName     string `mapstructure:"db.username" default:"postgres"`
	DBPassword     string `mapstructure:"db.password" default:"test@p4ssw0rd!"`
	DBDatabaseName string `mapstructure:"db.database" default:"qylo_drivers"`
	DBLogMode      int    `mapstructure:"db.logMode" default:"3"`
	DBLogLevel     int    `mapstructure:"db.logLevel" default:"3"`
	DBLogEnable    bool   `mapstructure:"db.logEnabled" default:"true"`
	DBLogThreshold int    `mapstructure:"db.logThreshold" default:"1"`

	JwtSecret string `mapstructure:"auth.secret-key" default:"secretw45h3re!"`

	RedisHost     string `mapstructure:"cache.configs.redis.host" default:"127.0.0.1"`
	RedisPort     string `mapstructure:"cache.configs.redis.port" default:"6380"`
	RedisPassword string `mapstructure:"cache.configs.redis.password" default:"OppoDev12@"`
}

func New() Config {
	cfg := Config{}

	viper.SetConfigName("app.config") // name of the config file (without extension)
	viper.AddConfigPath(".")          // path to look for the config file in
	viper.SetConfigType("yaml")       // type of the config file

	err := viper.ReadInConfig()
	if err != nil {
		return cfg
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg
	}

	return cfg
}
