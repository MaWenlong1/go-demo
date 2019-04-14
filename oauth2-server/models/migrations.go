package models

import (
	"mwl/oauth2-server/util/migrations"

	"github.com/jinzhu/gorm"
)

var (
	list = []migrations.MigrationStage{
		{
			Name:     "initial",
			Function: migrate001,
		},
	}
)

func migrate001(db *gorm.DB, name string) error {
	return nil
}
