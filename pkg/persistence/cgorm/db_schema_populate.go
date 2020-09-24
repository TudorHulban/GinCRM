package cgorm

import (
	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// PopulateSchemaSecurityRoles Inserts security rights and roles.
func PopulateSchemaSecurityRoles(l *log.LogInfo) error {
	var errSecu error // reuse var

	l.Debug("Populating persistence layer with security rights data:")

	for _, right := range setup.SecuRights {
		if errSecu = AddSecurityRight(right, l); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security right: %v", right)
		}
	}

	l.Debug("Populating persistence layer with security roles data:")

	for _, role := range setup.SecuRoles {
		if errSecu = AddSecurityRole(role, l); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security role: %v", role)
		}
	}

	l.Debug("Populating persistence layer with security roles definition:")

	for roleID, theRights := range setup.RolesDefinition {
		if errSecu = AddSecurityRoleDefinition(uint8(roleID), theRights, l); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security role: %v", roleID)
		}
	}

	return nil
}
