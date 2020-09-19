package cgorm

import (
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
)

// User Type concentrates user operations.
type User struct{}

// AddUser Method adds user to persistance.
func (User) AddUser(data *persistence.UserAuth) error {
	res := persistenceconn.GetRDBMSConn().Create(data)
	return res.Error
}
