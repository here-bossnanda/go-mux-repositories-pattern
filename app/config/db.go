package config

import (
	"database/sql"

	"gorm.io/gorm"
)

var (
	DBORM *gorm.DB
	DB    *sql.DB
	err   error
)
