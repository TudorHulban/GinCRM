package authentication

import (
	"github.com/TudorHulban/GinCRM/pkg/cache/cachelogin"
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/badgerwrap"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// OPAuthentication Structure defined for the authentication operation.
type OPAuthentication struct {
	useCache  bool
	data      UserAuth
	crudLogic persistence.IUserCRUD // access CRUD persistance layer
	l         *log.LogInfo
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
func NewOPAuthentication(credentials UserAuth, crud persistence.IUserCRUD, logger *log.LogInfo) *OPAuthentication {
	return &OPAuthentication{
		useCache:  true,
		data:      credentials,
		crudLogic: crud,
		l:         logger,
	}
}

// NewOPAuthenticationNoCache Calling authentication one would need a constructor
// for authentication operation.
func NewOPAuthenticationNoCache(credentials UserAuth, crud persistence.IUserCRUD, logger *log.LogInfo) *OPAuthentication {
	return &OPAuthentication{
		useCache:  false,
		data:      credentials,
		crudLogic: crud,
		l:         logger,
	}
}

// SaveToLoginCache Method saves user to login cache.
// TODO: based SOLID move to cache package.
func (op *OPAuthentication) SaveToLoginCache() error {
	return cachelogin.GetCache().Set(badgerwrap.KV{
		Key:   []byte(op.data.Code),
		Value: []byte(op.data.Password),
	})
}

// IsAuthenticated Method checks if credentials match the cache.
func (op *OPAuthentication) IsAuthenticated() error {
	// check if session ID exists and if yes the value in cache

	// check if using cache
	if op.useCache {
		errCache := op.isCachedAuthenticated()
		if errCache == nil {
			return nil
		}
	}

	// check if credentials persisted
	errPersisted := op.isPersistentAuthenticated()
	if errPersisted == nil {
		return nil
	}

	return ErrorUnknownCredentials
}

// isCachedAUthenticated Checks if app user is cached authenticated
func (op *OPAuthentication) isCachedAuthenticated() error {
	pass, errGet := cachelogin.GetCache().GetVByK([]byte(op.data.Code))
	if errGet != nil {
		return errors.WithMessage(errGet, "error when checking login cache")
	}

	if op.data.Password != string(pass) {
		return errors.New("password does not match login cache one")
	}

	return nil
}

// isPersistentAuthenticated Checks if user credentails are according to persisted values.
func (op *OPAuthentication) isPersistentAuthenticated() error {
	_, errGet := op.crudLogic.GetUserByCredentials(op.data.Code, op.data.Password)
	if errGet == nil {
		return nil
	}
	return ErrorUnknownCredentials
}

// DeleteFromCache Method deletes credentials from login cache.
func (op *OPAuthentication) DeleteFromCache(usercode string) error {
	return cachelogin.GetCache().DeleteKVByK([]byte(usercode))
}
