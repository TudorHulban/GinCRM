package cgorm

import (
	"time"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
	"github.com/pkg/errors"
)

// User Type concentrates user operations.
type User struct{}

// AddUser Method adds user to persistance.
func (User) AddUser(data *persistence.UserAuth) error {
	if errValid := validateStruct(data); errValid != nil {
		return errors.WithMessage(errValid, "validation error when adding user")
	}

	data.CreatedAt = time.Now().Unix()
	data.LastUpdateAt = time.Now().Unix()

	res := persistenceconn.GetRDBMSConn().Create(data)
	return res.Error
}
