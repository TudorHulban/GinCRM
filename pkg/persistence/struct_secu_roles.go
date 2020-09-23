package persistence

// SecurityRight Structure defined for persisting security rights.
type SecurityRight struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

// SecurityRole Structure defined for persisting security roles.
type SecurityRole struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

// RoleDefinition Structure defined for persisting role definition.
// A role is defined by the security rights it has.
type RoleDefinition struct {
	ID      uint64 `gorm:"primaryKey"`
	RoleID  uint8  `validate:"required"`
	RightID uint8  `validate:"required"`
}
