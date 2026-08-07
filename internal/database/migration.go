package database

import "gorm.io/gorm"

// RunMigrations applies schema migrations for the given models, converts any
// leftover soft-delete state to hard deletes, and then ensures the explicit
// foreign key constraints exist.
func RunMigrations(db *gorm.DB, values ...interface{}) error {
	if err := db.AutoMigrate(values...); err != nil {
		return err
	}

	return ApplyConstraints(db)
}
