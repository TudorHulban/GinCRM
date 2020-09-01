package httpinterface

import (
	"testing"

	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateConfig(t *testing.T) {
	const vers = "0.1"

	co, err := CreateConfig("0.0.0.0:80", vers, log.DEBUG, 3)

	assert.Nil(t, err)
	assert.Equal(t, co.IPV4Address, "0.0.0.0", "testing IPV4")
	assert.Equal(t, co.Port, uint16(80), "testing port")
	assert.Equal(t, co.BinaryVersion, vers)
}
