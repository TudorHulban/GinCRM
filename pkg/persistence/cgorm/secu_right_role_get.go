package cgorm

import (
	"fmt"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// SecurityRR Type concentrates security rights, roles and profiles operations.
type SecurityRR struct {
	l *log.LogInfo
}

// NewSecurityRR Constructor for user security related operations in persistence.
func NewSecurityRR(logger *log.LogInfo) persistence.ISecurityRoles {
	return &SecurityRR{
		l: logger,
	}
}

// GetSecurityRightsForRole Returns a slice with security rights given the role ID.
func (op *SecurityRR) GetSecurityRightsForRole(roleID uint8) ([]uint8, error) {
	op.l.Debugf("Fetching user rights for role ID:%v", roleID)

	var userRoleDef []persistence.SecurityDefRole
	persistenceconn.GetRDBMSConn().Where("role_id = ?", roleID).Find(&userRoleDef)

	if userRoleDef == nil {
		return nil, fmt.Errorf("no security roles found for role:%v", roleID)
	}

	result := make([]uint8, len(userRoleDef))
	for i, definition := range userRoleDef {
		result[i] = definition.RightID
	}

	op.l.Debugf("User rights for role ID:%v are:%v", roleID, result)
	return result, nil
}

// GetSecurityRightsForProfile Returns a slice with security rights given the profile ID.
func (op *SecurityRR) GetSecurityRightsForProfile(profileID uint8) ([]uint8, error) {
	op.l.Debugf("Fetching user rights for profile ID:%v", profileID)

	var userProfilesDef []persistence.SecurityDefProfile
	persistenceconn.GetRDBMSConn().Where("profile_id = ?", profileID).Find(&userProfilesDef)

	userRoles := make([]uint8, len(userProfilesDef))
	for i, definition := range userProfilesDef {
		userRoles[i] = definition.RoleID
	}

	var result []uint8
	for _, roleID := range userRoles {
		buf, errGet := op.GetSecurityRightsForRole(roleID)
		if errGet != nil {
			return nil, errors.WithMessagef(errGet, "when fetching security rights for role ID:%v", roleID)
		}
		op.l.Debugf("Appending for profile ID:%v following user rights:%v", profileID, buf)
		result = append(result, buf...)
	}

	return result, nil
}

func (op *SecurityRR) GetSecurityProfilesDefinition() (map[uint8][]uint8, error) {
	op.l.Debug("Fetching security profiles definition")

	var userProfiles []persistence.SecurityProfile
	persistenceconn.GetRDBMSConn().Find(&userProfiles)

	userProfiles := make([]uint8, len(userProfiles))

}
