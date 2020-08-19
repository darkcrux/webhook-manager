package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "run the server",
	RunE:  runServe,
}

func init() {
	root.AddCommand(serve)
}

func runServe(c *cobra.Command, args []string) error {

	log.Info("========================================")
	log.Info("Callback Manager")
	log.Info("========================================")

	if err := runMigrate(c, args); err != nil {
		log.WithError(err).Error("Migration failed")
		return err
	}

	// if err := setLogLevel(); err != nil {
	// 	log.WithError(err).Error("unable to set log level")
	// 	return err
	// }

	restAPI := createRestAPI()

	log.Info("========================================")
	log.Info("Starting API Server")
	log.Info("========================================")

	if err := restAPI.Run(); err != nil {
		log.WithError(err).Error("REST API terminated")
		return err
	}
	return nil
}

func setLogLevel() error {
	logLevel := viper.GetString("log.level")
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithError(err).Errorf("log level %s is invalid", logLevel)
		return err
	}
	log.SetLevel(level)
	log.Infof("Log Level Set to: %s", level)
	return nil
}
