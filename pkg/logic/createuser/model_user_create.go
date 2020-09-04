package createuser

import "github.com/TudorHulban/GinCRM/pkg/cache/cachesessionid"

/*
File contains creation only in KV.
RDBMS to be added.
*/

// UserFormData Concentrates information from front end for user creation.
type UserFormData struct {
	SecurityProfile uint8
	CompanyID       uint64
	TeamID          uint64
	FormUserCode    string
	FormPassword    string
}

// UserData Concentrates information for backend for user creation.
type UserData struct {
	ID int64
	UserFormData
}

// CreateUser Method to create user.
// Returns session ID and any error.
func (u UserData) CreateUser() (string, error) {
	// create in RDBMS
	// create session ID in cache
	// retrun session ID

	sessionID := generateSessionID()
	return sessionID, cachesessionid.GetCache().SetAny([]byte(sessionID), u)
}
