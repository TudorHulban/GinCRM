package cgorm

import (
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
)

// MigrateDBSchema Creates or updates database schema.
func MigrateDBSchema() error {
	tables := interface{}(&persistence.SecurityRole{})
	
	//[]interface{}{		interface{}(&persistence.SecurityRole{}),		interface{}(&persistence.RoleDefinition{}),,, persistence.SecurityRight{}
	} 



	return persistenceconn.GetRDBMSConn().AutoMigrate(persistence.User{})
}
