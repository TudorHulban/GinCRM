package authentication

/*
File details structure to return as result of authentication operation.
Provides only one method which is for logout.
*/

// UserAuthInfo Concentrates response data for authentication operations.
type UserAuthInfo struct {
	UserID    int64 // primary key as per user table.
	UserCode  string
	UserRoles []uint // A security profile is a list of roles.
}

// Logout Performs clean up from credentials and session cache
func (a UserAuthInfo) Logout() error {
	logoutCredentialsCache(a.UserCode)
}
