package persistenceconn_test

import (
	"testing"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
)

type TableA struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

type TableB struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

func TestTablesCreation(t *testing.T) {
	setup.CleanRDBMS()

	db := persistenceconn.GetRDBMSConn()
	db.
}
