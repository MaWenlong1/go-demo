package cmd

import (
	"mwl/gravitee/config"
	"mwl/gravitee/database"

	"github.com/jinzhu/gorm"
)

// initConfigDB loads the configuration and connects to the database
func initConfigDB(configFile string) (*config.Config, *gorm.DB, error) {
	// Config
	cfg := config.NewConfig(configFile)

	// Databse
	db, err := database.NewDatabase(cfg)
	if err != nil {
		return nil, nil, err
	}

	return cfg, db, nil
}
