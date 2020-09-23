package cgorm

import (
	"github.com/pkg/errors"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
	"github.com/TudorHulban/log"
)

// AddSecurityRight Helper for adding security right.
func AddSecurityRight(description string, logger *log.LogInfo) error {
	data := persistence.SecurityRight{
		Description: description,
	}
	if errValid := validateStruct(&data); errValid != nil {
		return errors.WithMessage(errValid, "validation error when adding security right")
	}
	res := persistenceconn.GetRDBMSConn().Create(&data)
	return res.Error
}

// AddSecurityRole Helper for adding security role.
func AddSecurityRole(description string, logger *log.LogInfo) error {
	data := persistence.SecurityRole{
		Description: description,
	}
	if errValid := validateStruct(&data); errValid != nil {
		return errors.WithMessage(errValid, "validation error when adding security role")
	}
	res := persistenceconn.GetRDBMSConn().Create(&data)
	return res.Error
}
