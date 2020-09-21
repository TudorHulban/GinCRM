package cgorm

import (
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
)

// GetUserByCredentials Fetches user by credentials.
func (u *User) GetUserByCredentials(userCode, password string) (*persistence.User, error) {
	u.l.Debugf("Fetching user data by credential user:%v, password:%v", userCode, password)

	var fetchedUserData persistence.User
	res := persistenceconn.GetRDBMSConn().Where(&persistence.User{UserCode: userCode}).First(&fetchedUserData)
	if res.Error != nil {
		return nil, ErrorDatabase
	}

	return &fetchedUserData, nil
}
