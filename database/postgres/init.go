package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	driverPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type psql struct {
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
type pgsqlOption func(*psql)

func Connect(DBHost string, DBPort int, DBUserName string, DBPassword string, DBDatabaseName string, options ...pgsqlOption) (*gorm.DB, error) {
	db := &psql{
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

func connect(param *psql) (*gorm.DB, error) {
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

	db, err := gorm.Open(driverPostgres.Open(dsn), cfg)

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

func SetMaxIdleConns(conns int) pgsqlOption {
	return func(c *psql) {
		if conns > 0 {
			c.maxIdleConnection = conns
		}
	}
}

func SetMaxOpenConns(conns int) pgsqlOption {
	return func(c *psql) {
		if conns > 0 {
			c.maxOpenConnection = conns
		}
	}
}

func SetConnMaxLifetime(conns int) pgsqlOption {
	return func(c *psql) {
		if conns > 0 {
			c.connectionMaxLifetimeInSecond = conns
		}
	}
}

func SetNamingStrategy(namingStrategy schema.Namer) pgsqlOption {
	return func(c *psql) {
		c.namingStrategy = namingStrategy
	}
}

// SetPrintLog level: 1=silent, 2=Error, 3=Warn, 4=Info. latencyThreshold: suggestion 200ms.
func SetPrintLog(isenable bool, level logger.LogLevel, latencyThreshold time.Duration) pgsqlOption {
	return func(c *psql) {
		if latencyThreshold > 0 {
			c.printLog = isenable
			c.logLevel = level
			c.logThreshold = latencyThreshold
		}
	}
}

func SetTimezone(timezone string) pgsqlOption {
	return func(c *psql) {
		if timezone == "" {
			c.DBTimezone = "Etc/UTC"
		} else {
			c.DBTimezone = timezone
		}
	}
}
