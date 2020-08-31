package httpinterface

import (
	"github.com/TudorHulban/log"
	"github.com/gin-gonic/gin"
)

// Config Concentrates attributes for starting a Gin server.
type Config struct {
	// IPV4 to run
	IPV4Address string
	// port to run
	Port uint16
	// binary version
	BinaryVersion string
	// logger to use with Gin
	GLogger *log.LogInfo
}

// HTTPServer is HTTP server wrapper.
type HTTPServer struct {
	Config
	isReady func() bool // for k8s probes
	engine  *gin.Engine
}

// NewGinServer creates a new HTTP server
func NewGinServer(cfg Config) *HTTPServer {
	return &HTTPServer{
		engine: gin.Default(),
	}
}
