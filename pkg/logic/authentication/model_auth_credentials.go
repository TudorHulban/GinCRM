package authentication

import (
	"github.com/TudorHulban/GinCRM/pkg/cache/cachecredentials"
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/badgerwrap"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// Credentials Auth Concentrates user authentication data based on user and passowrd credentials.
type Credentials struct {
	Code     string
	Password string
}

// OPAuthenticationCredentials Structure defined for the authentication operations.
type OPAuthenticationCredentials struct {
	useCache  bool
	data      Credentials
	crudLogic persistence.IUserCRUD // access CRUD persistance layer
	l         *log.LogInfo
}

// NewOPAuthenticationCredentials Calling authentication one would need a constructor
// for authentication operation.
func NewOPAuthenticationCredentials(credentials Credentials, crud persistence.IUserCRUD, logger *log.LogInfo) *OPAuthenticationCredentials {
	return &OPAuthenticationCredentials{
		useCache:  true,
		data:      credentials,
		crudLogic: crud,
		l:         logger,
	}
}

// NewOPAuthenticationCredentialsNoCache Calling authentication one would need a constructor
// for authentication operation.
func NewOPAuthenticationCredentialsNoCache(credentials Credentials, crud persistence.IUserCRUD, logger *log.LogInfo) *OPAuthenticationCredentials {
	return &OPAuthenticationCredentials{
		useCache:  false,
		data:      credentials,
		crudLogic: crud,
		l:         logger,
	}
}

// CanLogin Method checks if credentials match the cache or persisted.
func (op *OPAuthenticationCredentials) CanLogin() error {
	// check if credentials exist in cache
	if op.useCache {
		if errCache := op.isCachedAuthenticated(); errCache == nil {
			return nil
		}
	}

	// check if credentials persisted
	if errPersisted := op.isPersistentAuthenticated(); errPersisted == nil {
		return nil
	}

	return ErrorUnknownCredentials
}

// isCachedAUthenticated Checks if app user is cached authenticated
func (op *OPAuthenticationCredentials) isCachedAuthenticated() error {
	pass, errGet := cachecredentials.GetCache().GetVByK([]byte(op.data.Code))
	if errGet != nil {
		return errors.WithMessage(errGet, "error when checking login cache")
	}

	if op.data.Password != string(pass) {
		return errors.New("password does not match login cache one")
	}

	return nil
}

// isPersistentAuthenticated Checks if user credentails are according to persisted values.
func (op *OPAuthenticationCredentials) isPersistentAuthenticated() error {
	userData, errGet := op.crudLogic.GetUserByCredentials(op.data.Code, op.data.Password)
	if errGet != nil {
		return errGet
	}
	if userData == nil {
		op.l.Debugf("No user with user/passwd: %v/%v", op.data.Code, op.data.Password)
		return ErrorUnknownCredentials
	}
	if checkPasswordHash(op.data.Password, userData.PasswordSALT, userData.PasswordHASH) {
		return nil
	}
	op.l.Debugf("Bad password: %v for user: %v", op.data.Password, userData.ID)
	return ErrorUnknownCredentials
}

// saveToLoginCache Method saves credentials to login cache.
func (op *OPAuthenticationCredentials) saveToCredentialsCache() error {
	return cachecredentials.GetCache().Set(badgerwrap.KV{
		Key:   []byte(op.data.Code),
		Value: []byte(op.data.Password),
	})
}
