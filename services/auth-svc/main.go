package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm/logger"

	"go-monorepo-boilerplate/database/postgres"
	userUc "go-monorepo-boilerplate/services/auth-svc/app/usecase/user"
	"go-monorepo-boilerplate/services/auth-svc/config"
	"go-monorepo-boilerplate/services/auth-svc/handler"
	"go-monorepo-boilerplate/services/auth-svc/repository/user"
)

// @title OpenAPI - Swagger Documentation
// @version 1.0
// @description Authentication Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Go Monorepo
// @contact.url https://github.com/krisnasw
// @contact.email hello@sociolite.id

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 0.0.0.0:29000
// @BasePath /api/v1
// @schemes http
func main() {
	config.New()

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-monorepo"),
		newrelic.ConfigLicense("6e9314f0c0ff95dde04848f40734d10fcdb6NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	router := gin.Default()
	router.Use(nrgin.Middleware(app))
	router.Use(nrgin.MiddlewareHandlerTxnNames(app))
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	db, err := postgres.Connect(viper.GetString("db.host"), viper.GetInt("db.port"), viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.database"),
		postgres.SetPrintLog(viper.GetBool("db.logEnabled"), logger.LogLevel(viper.GetInt("db.logLevel")), time.Duration(viper.GetInt("db.logThreshold"))*time.Millisecond))
	if err != nil {
		logrus.Errorln("Failed to initialized connection postgres DB:", err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              "https://1b9d15a3b6334f098d0258676db9f22f@o4504836237950976.ingest.sentry.io/4504857583943680",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	userRepository := user.New(db, "users")
	userUseCase := userUc.NewUseCase(userRepository)

	h := handler.New(userUseCase)

	api := router.Group("/api/v1")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.POST("/login", h.Login)
	}

	router.Run(fmt.Sprintf(":%d", viper.GetInt("services.auth-service.port")))
}
