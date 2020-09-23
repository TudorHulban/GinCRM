package cgorm

import (
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
)

// MigrateDBSchema Creates or updates database schema.
func MigrateDBSchema() error {
	tables := []interface{}{
		interface{}(&persistence.User{}),
		interface{}(&persistence.SecurityRight{}),
		interface{}(&persistence.SecurityRole{}),
		interface{}(&persistence.RoleDefinition{}),
	}
	return persistenceconn.GetRDBMSConn().AutoMigrate(tables)
}
