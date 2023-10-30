package repositories

import (
	"load-balancer/database"

	"gorm.io/gorm"
)

var DB *gorm.DB = database.Init()
