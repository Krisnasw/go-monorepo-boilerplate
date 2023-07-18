package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	driverMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type mysql struct {
	DBHost         string
	DBPort         int
	DBUserName     string
	DBPassword     string
	DBDatabaseName string
	DBTimezone     string

	printLog     bool
	logLevel     logger.LogLevel
	logThreshold time.Duration

	maxIdleConnection             int
	maxOpenConnection             int
	connectionMaxLifetimeInSecond int
	namingStrategy                schema.Namer
}
type mysqlOption func(*mysql)

func Connect(DBHost string, DBPort int, DBUserName string, DBPassword string, DBDatabaseName string, options ...mysqlOption) (*gorm.DB, error) {
	db := &mysql{
		DBHost:         DBHost,
		DBPort:         DBPort,
		DBUserName:     DBUserName,
		DBPassword:     DBPassword,
		DBDatabaseName: DBDatabaseName,
		DBTimezone:     "Etc/UTC",

		printLog:     false,
		logLevel:     logger.Silent,
		logThreshold: 200 * time.Millisecond,

		maxIdleConnection:             5,
		maxOpenConnection:             10,
		connectionMaxLifetimeInSecond: 60,
		namingStrategy:                nil,
	}

	for _, o := range options {
		o(db)
	}

	return connect(db)
}

func connect(param *mysql) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		param.DBHost, param.DBUserName, param.DBPassword, param.DBDatabaseName, param.DBPort, param.DBTimezone)

	cfg := &gorm.Config{}
	if param.printLog {
		cfg.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: param.logThreshold,
			LogLevel:      param.logLevel,
			Colorful:      true,
		})
	}
	if param.namingStrategy != nil {
		cfg.NamingStrategy = param.namingStrategy
	}

	db, err := gorm.Open(driverMysql.Open(dsn), cfg)

	if err != nil {
		return nil, err
	}
	mysqlDb, err := db.DB()
	mysqlDb.SetMaxOpenConns(param.maxOpenConnection)
	mysqlDb.SetConnMaxLifetime(time.Duration(param.connectionMaxLifetimeInSecond) * time.Minute)
	mysqlDb.SetMaxIdleConns(param.maxOpenConnection)
	if err != nil {
		return nil, err
	}
	return db, err
}

func SetMaxIdleConns(conns int) mysqlOption {
	return func(c *mysql) {
		if conns > 0 {
			c.maxIdleConnection = conns
		}
	}
}

func SetMaxOpenConns(conns int) mysqlOption {
	return func(c *mysql) {
		if conns > 0 {
			c.maxOpenConnection = conns
		}
	}
}

func SetConnMaxLifetime(conns int) mysqlOption {
	return func(c *mysql) {
		if conns > 0 {
			c.connectionMaxLifetimeInSecond = conns
		}
	}
}

func SetNamingStrategy(namingStrategy schema.Namer) mysqlOption {
	return func(c *mysql) {
		c.namingStrategy = namingStrategy
	}
}

// SetPrintLog level: 1=silent, 2=Error, 3=Warn, 4=Info. latencyThreshold: suggestion 200ms.
func SetPrintLog(isenable bool, level logger.LogLevel, latencyThreshold time.Duration) mysqlOption {
	return func(c *mysql) {
		if latencyThreshold > 0 {
			c.printLog = isenable
			c.logLevel = level
			c.logThreshold = latencyThreshold
		}
	}
}

func SetTimezone(timezone string) mysqlOption {
	return func(c *mysql) {
		if timezone == "" {
			c.DBTimezone = "Etc/UTC"
		} else {
			c.DBTimezone = timezone
		}
	}
}
