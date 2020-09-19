package cgorm_test

import (
	"testing"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	var userCRUD cgorm.User

	data := persistence.UserAuth{
		SecurityGroupID: 1,
		UserCode:        "abcd",
	}

	assert.Nil(t, cgorm.MigrateDBSchema())
	assert.Nil(t, userCRUD.AddUser(&data))
}
