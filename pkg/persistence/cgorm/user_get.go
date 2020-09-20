package cgorm

import (
	"github.com/TudorHulban/GinCRM/pkg/logic/authentication"
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
)

// GetUserByCredentials Fetches user by credentials.
func (u *User) GetUserByCredentials(userCode, password string) (*persistence.User, error) {
	u.l.Debug("Fetching user data by credentials:", userCode, password)

	var fetchedUserData persistence.User
	res := persistenceconn.GetRDBMSConn().Where(&persistence.User{UserCode: userCode}).First(&fetchedUserData)
	if res.Error != nil {
		return nil, ErrorDatabase
	}

	// compare hashed password
	if password != fetchedUserData.PasswordHASH {
		return nil, authentication.ErrorUnknownCredentials
	}

	return &fetchedUserData, nil
}
