package httpinterface

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	cfg     *Config
	isReady func() bool // for probes
	engine  *gin.Engine
}

// NewGinServer creates a new HTTP server.
// No cfg validation, use create config helper for that.
func NewGinServer(config *Config) *HTTPServer {
	s := new(HTTPServer)
	s.SetAsNotReady()
	s.cfg = config

	if s.cfg.GLogger.GetLogLevel() != log.DEBUG {
		s.cfg.GLogger.Info("Setting Gin Log Level to Release Mode")
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

// Run Method for starting HTTP Server.
func (s *HTTPServer) Run(ctx context.Context) error {
	gracefull := &http.Server{
		Handler: s.engine,
		Addr:    fmt.Sprintf("%d:%d", s.cfg.IPV4Address, s.cfg.Port),
	}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		s.cfg.GLogger.Info("shutting down HTTP Server")
		if errShutdown := gracefull.Shutdown(ctx); errShutdown != nil {
			s.cfg.GLogger.Info(errShutdown, "could not gracefully stop http server")
		}
	}()

	return gracefull.ListenAndServe()
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
