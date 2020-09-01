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
	isReady func() bool // for probes
	engine  *gin.Engine
}

// NewGinServer creates a new HTTP server.
// No cfg validation, use create config helper for that.
func NewGinServer(cfg Config) *HTTPServer {
	s := new(HTTPServer)
	s.SetAsNotReady()
	s.Config = cfg

	if s.GLogger.GetLogLevel() != log.DEBUG {
		s.GLogger.Info("Setting Gin Log Level to Release Mode")
		gin.SetMode(gin.ReleaseMode) // to be called before initializing the router!
	}
	s.engine = gin.New()
	s.engine.Use(gin.Recovery())
	s.engine.RedirectTrailingSlash = true
	s.engine.HandleMethodNotAllowed = false

	// adding routes
	s.registerRoutes(s.prepareInfraRoutes())

	return s
}

// SetAsReady Setter method for switching to readiness state READY.
func (s *HTTPServer) SetAsReady() {
	s.isReady = func() bool { return true }
}

// SetAsNotReady Setter method for switching to readiness state NOTREADY.
func (s *HTTPServer) SetAsNotReady() {
	s.isReady = func() bool { return false }
}

// GetIfReady Getter method for current readiness state.
func (s *HTTPServer) GetIfReady() bool {
	return s.isReady()
}
