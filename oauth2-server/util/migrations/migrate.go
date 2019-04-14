package migrations

import (
	"github.com/jinzhu/gorm"
)

// MigrationStage ..
type MigrationStage struct {
	Name     string
	Function func(db *gorm.DB, name string) error
}
