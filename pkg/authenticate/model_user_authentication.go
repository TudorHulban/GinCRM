package autentication

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

// SaveToCache Method saves user to login cache.
func (u UserAuth) SaveToCache() error {
	return cachelogin.GetCache().Set(badgerwrap.KV{
		Key:   []byte(u.Code),
		Value: []byte(u.Password),
	})
}

// IsAuthenticated Method checks if credentials match the cache.
func (u UserAuth) IsAuthenticated(usercode, password string) error {
	pass, errGet := cachelogin.GetCache().GetVByK([]byte(usercode))
	if errGet != nil {
		return errors.WithMessage(errGet, "error when checking login cache")
	}

	if password != string(pass) {
		return errors.New("password does not match login cache one")
	}

	return nil
}

// DeleteFromCache Method deletes credentials from login cache.
func (u UserAuth) DeleteFromCache(usercode string) error {
	return cachelogin.GetCache().DeleteKVByK([]byte(usercode))
}
