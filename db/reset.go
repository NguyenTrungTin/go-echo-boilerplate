package db

import "github.com/jinzhu/gorm"

// DestructiveReset is used to delete all data and table on the database
func DestructiveReset() *gorm.DB {
	return db.DropTableIfExists(models...)
}
