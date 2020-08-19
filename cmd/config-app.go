package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	appConfig struct {
		Log        logConfig        `mapstructure:"log"`
		API        apiConfig        `mapstructure:"api"`
		Datasource datasourceConfig `mapstructure:"datasource"`
	}

	apiConfig struct {
		REST restConfig `mapstructure:"rest"`
	}

	restConfig struct {
		Host     string         `mapstructure:"host"`
		Port     int            `mapstructure:"port"`
		Spec     string         `mapstructure:"spec"`
		CORS     corsConfig     `mapstructure:"cors"`
		Security securityConfig `mapstructure:"security"`
	}

	corsConfig struct {
		AllowedOrigins []string `mapstructure:"allowedOrigins"`
		AllowedHeaders []string `mapstructure:"allowedHeaders"`
		AllowedMethods []string `mapstructure:"allowedMethods"`
	}

	securityConfig struct {
		JWTPublicKey string `mapstructure:"jwtPublicKey"`
		RBACConfig   string `mapstructure:"rbacConfig"`
	}

	datasourceConfig struct {
		Type       string `mapstructure:"type"`
		Host       string `mapstructure:"host"`
		Port       int    `mapstructure:"port"`
		Database   string `mapstructure:"database"`
		SSLMode    string `mapstructure:"sslMode"`
		Migrations string `mapstructure:"migrations"`
		Username   string `mapstructure:"username"`
		Password   string `mapstructure:"password"`
	}

	secretsConfig struct {
		File string `mapstructure:"file"`
	}

	logConfig struct {
		Level string `mapstructure:"level"`
	}
)

func defaultConfig() appConfig {
	return appConfig{
		Log: logConfig{
			Level: "info",
		},
		API: apiConfig{
			REST: restConfig{
				Host: "0.0.0.0",
				Port: 8080,
				Spec: "./openapi.yaml",
				CORS: corsConfig{
					AllowedOrigins: []string{"*"},
					AllowedHeaders: []string{
						"Content-Type",
						"Sec-Fetch-Dest",
						"Referer",
						"accept",
						"Sec-Fetch-Mode",
						"Sec-Fetch-Site",
						"User-Agent",
						"User-Agent",
						"API-KEY",
						"Authorization",
					},
					AllowedMethods: []string{
						"OPTIONS",
						"GET",
						"POST",
						"DELETE",
					},
				},
			},
		},
		Datasource: datasourceConfig{
			Type:       "postgres",
			Host:       "localhost",
			Port:       5432,
			Database:   "webhook_manager",
			SSLMode:    "disable",
			Migrations: "db/migrations",
			Username:   "user",
			Password:   "password",
		},
	}
}

func loadConfig() (config appConfig, err error) {
	log.Info("Loading App Config...")
	config = defaultConfig()
	err = viper.Unmarshal(&config)
	if err != nil {
		log.WithError(err).Error("unable to load app config")
	}
	return
}
