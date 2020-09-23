package cgorm

import (
	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// PopulateSchemaSecurityRoles Inserts security rights and roles.
func PopulateSchemaSecurityRoles(logger *log.LogInfo) error {
	var errSecu error // reuse var

	for _, right := range setup.SecuRights {
		if errSecu = AddSecurityRight(right, logger); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security right: %v", right)
		}
	}

	for _, role := range setup.SecuRoles {
		if errSecu = AddSecurityRole(role, logger); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security role: %v", role)
		}
	}

	return nil
}
