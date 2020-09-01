package httpinterface

import (
	"net/http"
	"testing"

	"github.com/TudorHulban/log"
	"github.com/steinfletcher/apitest"
)

func TestHandlersInfrastructure(t *testing.T) {
	const socket = "0.0.0.0:8001"
	cfg, _ := CreateConfig(socket, "0.2.0", log.DEBUG, 1)

	s := NewGinServer(cfg)

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
