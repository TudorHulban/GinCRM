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
			return errors.WithMessage(errValid, " adding security right")
		}
	}

	for _, role := range setup.SecuRoles {
		AddSecurityRight(role, logger)
	}
}
