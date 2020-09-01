package httpinterface

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handlerGetIfReady Method is infra handler that returns HTTP status 200 and JSON with if app ready.
func (s *HTTPServer) handlerGetIfReady(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"isReady": s.GetIfReady()})
}

// handlerVersion Method returns binary version as injected by repo when created the Gin instance.
func (s *HTTPServer) handlerVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": s.Config.BinaryVersion})
}

// prepareInfraRoutes Method helps with route preparation.
// Routes need to contain the starting slash ex. /route.
func (s *HTTPServer) prepareInfraRoutes() []route {
	routeIsReady := route{
		Group:    endPointGroupInfrastructure,
		Endpoint: endpointIsReady,
		Method:   "GET",
		Handler:  s.handlerGetIfReady,
	}

	routeVersion := route{
		Group:    endPointGroupInfrastructure,
		Endpoint: endpointVersion,
		Method:   "GET",
		Handler:  s.handlerVersion,
	}

	return []route{routeIsReady, routeVersion}
}
