package httpinterface

import (
	"net/http"
	"testing"

	"github.com/TudorHulban/log"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func TestHandlersInfrastructure(t *testing.T) {
	const socket = "0.0.0.0:8001"
	cfg, _ := CreateConfig(socket, "0.2.0", log.DEBUG, 1)

	s, errCo := NewGinServer(cfg)
	if assert.Nil(t, errCo) {
		apitest.New().
			Handler(s.engine).
			Get(endPointGroupInfrastructure + endpointIsReady).
			Expect(t).
			Status(http.StatusOK).
			End()

		apitest.New().
			Handler(s.engine).
			Get(endPointGroupInfrastructure + endpointVersion).
			Expect(t).
			Status(http.StatusOK).
			End()
	}
}
