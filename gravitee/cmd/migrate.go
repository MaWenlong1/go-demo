package cmd

import (
	"mwl/gravitee/models"
	"mwl/gravitee/util/migrations"
)

// Migrate runs database migrations
func Migrate(configFile string) error {
	_, db, err := initConfigDB(configFile)
	if err != nil {
		return err
	}
	defer db.Close()

	// Bootstrap migrations
	if err := migrations.Bootstrap(db); err != nil {
		return err
	}

	// Run migrations for the oauth service
	if err := models.MigrateAll(db); err != nil {
		return err
	}

	return nil
}
