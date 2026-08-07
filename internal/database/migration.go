package database

import "gorm.io/gorm"

// RunMigrations applies schema migrations for the given models.
// Models are registered incrementally, so callers pass the entities
// that should be migrated.
func RunMigrations(db *gorm.DB, values ...interface{}) error {
	return db.AutoMigrate(values...)
}