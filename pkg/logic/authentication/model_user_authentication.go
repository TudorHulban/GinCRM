package authentication

import (
	"github.com/TudorHulban/GinCRM/pkg/cache/cachelogin"
	"github.com/TudorHulban/badgerwrap"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// OPAuthentication Structure defined for the authentication operation.
type OPAuthentication struct {
	l *log.LogInfo
}

// UserAuth Concentrates user authentication based on password or session ID.
// No need for constructor, passing by value.
type UserAuth struct {
	Code      string
	Password  string
	SessionID string // if session ID exists do not check by passowrd
}

// UserSessionInfo Concentrates info for when retrieving user info bassed on passed session ID.
// TODO: refactor.
type UserSessionInfo struct {
	UserID          int64 // Primary key as per user table.
	Code            string
	PermissionRoles []uint // A security profile is a list of roles.
}

// NewOPAuthentication Calling authentication one would need a constructor
// for authentication operation.
func NewOPAuthentication(credentials UserAuth, logger *log.LogInfo) *OPAuthentication {
	return &OPAuthentication{
		l: logger,
	}
}

// SaveToLoginCache Method saves user to login cache.
func (u UserAuth) SaveToLoginCache() error {
	return cachelogin.GetCache().Set(badgerwrap.KV{
		Key:   []byte(u.Code),
		Value: []byte(u.Password),
	})
}

// IsAuthenticated Method checks if credentials match the cache.
func (u UserAuth) IsAuthenticated() error {
	// check if session ID exists and if yes the value in cache

	// check if credentials in cache first
	errCache := u.isCachedAuthenticated()
	if errCache == nil {
		return nil
	}

	// check if credentials persisted

	return ErrorUnknownCredentials
}

// isCachedAUthenticated Checks if app user is cached authenticated
func (u UserAuth) isCachedAuthenticated() error {
	pass, errGet := cachelogin.GetCache().GetVByK([]byte(u.Code))
	if errGet != nil {
		return errors.WithMessage(errGet, "error when checking login cache")
	}

	if u.Password != string(pass) {
		return errors.New("password does not match login cache one")
	}

	return nil
}

// isPersistentAuthenticated Checks if user credentails are according to persisted values.
func (u UserAuth) isPersistentAuthenticated() error {
	return nil
}

// DeleteFromCache Method deletes credentials from login cache.
func (u UserAuth) DeleteFromCache(usercode string) error {
	return cachelogin.GetCache().DeleteKVByK([]byte(usercode))
}
