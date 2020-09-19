package cgorm

import (
	"time"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// User Type concentrates user operations.
type User struct {
	l *log.LogInfo
}

// NewUser Constructor for user related operations in persistence.
func NewUser(logger *log.LogInfo) persistence.IUserCRUD {
	return &User{
		l: logger,
	}
}

// AddUser Method adds user to persistance.
func (u *User) AddUser(data *persistence.UserAuth) error {
	u.l.Debug("Adding user:", data)

	if errValid := validateStruct(data); errValid != nil {
		return errors.WithMessage(errValid, "validation error when adding user")
	}

	data.CreatedAt = time.Now().Unix()
	data.LastUpdateAt = time.Now().Unix()

	res := persistenceconn.GetRDBMSConn().Create(data)
	return res.Error
}
