package Data

import (
	"YZT/Config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(mysql.Open(Config.DSN), &gorm.Config{})
