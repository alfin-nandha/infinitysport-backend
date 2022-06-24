package migration

import (
	_mUsers "project/e-comerce/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUsers.User{})
}
