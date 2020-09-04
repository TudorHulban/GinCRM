package authentication

import (
	"github.com/TudorHulban/GinCRM/pkg/cache/cachelogin"
	"github.com/TudorHulban/badgerwrap"
	"github.com/pkg/errors"
)

// UserAuth Concentrates user authentication based on password.
type UserAuth struct {
	Code     string
	Password string
}

// UserSessionInfo Concentrates info for when retrieving user info bassed on passed session ID.
type UserSessionInfo struct {
	UserID          int64 // Primary key as per user table.
	Code            string
	PermissionRoles []uint // A security profile is a list of roles.
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
	pass, errGet := cachelogin.GetCache().GetVByK([]byte(u.Code))
	if errGet != nil {
		return errors.WithMessage(errGet, "error when checking login cache")
	}

	if u.Password != string(pass) {
		return errors.New("password does not match login cache one")
	}

	return nil
}

// DeleteFromCache Method deletes credentials from login cache.
func (u UserAuth) DeleteFromCache(usercode string) error {
	return cachelogin.GetCache().DeleteKVByK([]byte(usercode))
}
