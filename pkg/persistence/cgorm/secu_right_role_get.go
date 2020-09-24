package cgorm

import (
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/log"
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

// GetSecurityRightsForProfile Returns a slice with security rights given the profile ID.
func (op *SecurityRR) GetSecurityRightsForProfile(profileID uint8) ([]uint8, error) {
	op.l.Debugf("Fetching user rights for profile ID:%v", profileID)

}
