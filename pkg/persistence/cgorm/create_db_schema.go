package cgorm

import (
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
)

// MigrateDBSchema Creates or updates database schema.
func MigrateDBSchema() error {
	return persistenceconn.GetRDBMSConn().AutoMigrate(interface{}(&persistence.UserAuth{}))
}