package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "run the db migrations",
	RunE:  runMigrate,
}

func init() {
	root.AddCommand(migrate)
}

func runMigrate(c *cobra.Command, args []string) error {
	migration := createMigration()
	if err := migration.Run(); err != nil {
		log.WithError(err).Error("DB Migration Error")
		return err
	}
	return nil
}
