package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	fileConfigName = "app.config"
	fileConfigPath = "../.."
	fileConfigType = "yml"
)

func New() {
	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigType(fileConfigType)
	viper.SetConfigName(fileConfigName)

	setDefaultKeys()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}

	logrus.Infof("initialized configs viper: success", fileConfigPath+"/"+fileConfigName+"."+fileConfigType)

	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	logrus.Infof("initialized WatchConfig(): success", fileConfigPath+"/"+fileConfigName+"."+fileConfigType)

	return
}

func setDefaultKeys() {
	viper.SetDefault("application.port", 29000)
	viper.SetDefault("application.mode", "debug")

	host := []string{"localhost", "https://gin-gonic.com"}
	viper.SetDefault("application.cors.allowedHost", host)

	//viper.SetDefault("db.configs.username", "root")
	//viper.SetDefault("db.configs.password", "password")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.database", "echo_sample")
	viper.SetDefault("db.maxIdleConn", 5)
	viper.SetDefault("db.maxOpenConn", 10)
	viper.SetDefault("db.logEnabled", true)
	viper.SetDefault("db.logMode", 3)
	viper.SetDefault("db.logLevel", 3)
	viper.SetDefault("db.logThreshold", true)

	// viper.SetDefault("cache.configs.redis.username", "root")
	// viper.SetDefault("cache.configs.redis.password", "password")
	viper.SetDefault("cache.configs.redis.db", 0)
	viper.SetDefault("cache.configs.redis.poolSize", 10)

	viper.SetDefault("cache.configs.redis.host", "127.0.0.1")
	viper.SetDefault("cache.configs.redis.port", 6379)

	viper.SetDefault("cache.ttl.short-period", "3h")
	viper.SetDefault("cache.ttl.medium-period", "24h")
	viper.SetDefault("cache.ttl.long-period", "3d")

	logrus.Infof("initialized default configs value : success", viper.AllSettings())
}
