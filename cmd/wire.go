//+build wireinject

package cmd

import (
	"os"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	//components
	"github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest"
	"github.com/darkcrux/webhook-manager/internal/infrastructure/postgres"

	repository "github.com/darkcrux/webhook-manager/internal/infrastructure/postgres/repository"

	txtypesService "github.com/darkcrux/webhook-manager/internal/component/txtypes"
	txtypesHandler "github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest/txtypes"

	customerService "github.com/darkcrux/webhook-manager/internal/component/customer"
	customerHandler "github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest/customer"

	webhookService "github.com/darkcrux/webhook-manager/internal/component/webhook"
	webhookHandler "github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest/webhook"
)

func createRestAPI() *rest.API {
	wire.Build(
		ProvideConfig,
		ProvideDatasource,
		ProvideGormDB,

		// txtypes
		repository.NewGormTxTypeRepository,
		txtypesService.NewDefaultService,
		txtypesHandler.NewController,

		// customer
		repository.NewGormCustomerRepository,
		customerService.NewDefaultService,
		customerHandler.NewController,

		// webhook
		repository.NewGormWebhookRepository,
		webhookService.NewDefaultService,
		webhookHandler.NewController,

		mux.NewRouter,
		ProvideRestAPIConfig,
		rest.NewRestAPI,
	)
	return &rest.API{}
}

func createMigration() *postgres.Migration {
	wire.Build(
		ProvideConfig,
		ProvideDatasource,
		postgres.NewMigration,
	)
	return &postgres.Migration{}
}

func ProvideConfig() appConfig {
	config, err := loadConfig()
	if err != nil {
		log.WithError(err).Error("unable to unmarshal configuration")
		os.Exit(1)
	}
	return config
}

func ProvideDatasource(config appConfig) *postgres.Datasource {
	return &postgres.Datasource{
		Type:       config.Datasource.Type,
		Host:       config.Datasource.Host,
		Port:       config.Datasource.Port,
		Database:   config.Datasource.Database,
		Username:   config.Datasource.Username,
		Password:   config.Datasource.Password,
		SSLMode:    config.Datasource.SSLMode,
		Migrations: config.Datasource.Migrations,
	}
}

func ProvideRestAPIConfig(config appConfig) *rest.Config {
	restConfig := &rest.Config{
		Host:    config.API.REST.Host,
		Port:    config.API.REST.Port,
		Spec:    config.API.REST.Spec,
		Version: "dev", // get this from input
		Cors: rest.CORSConfig{
			AllowedOrigins: config.API.REST.CORS.AllowedOrigins,
			AllowedHeaders: config.API.REST.CORS.AllowedHeaders,
			AllowedMethods: config.API.REST.CORS.AllowedMethods,
		},
	}
	log.Info("========================================")
	log.Info("API Configuration")
	log.Info("========================================")
	log.Info("Host:    ", restConfig.Host)
	log.Info("Port:    ", restConfig.Port)
	log.Info("Spec:    ", restConfig.Spec)
	log.Info("Version: ", restConfig.Port)
	return restConfig
}

func ProvideGormDB(datasource *postgres.Datasource) *gorm.DB {
	db, err := gorm.Open("postgres", datasource.AsPQString())
	if err != nil {
		log.WithError(err).Error("unable to get gorm db connection")
		os.Exit(1)
	}
	return db
}
