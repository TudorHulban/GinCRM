package httpinterface

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// FormLogin Structure used for validating login request.
type FormLogin struct {
	FieldUserCode string `form:"usercode"`
	FieldPassword string `form:"password"`
}

func (s *HTTPServer) handlerLogin(c *gin.Context) {
	var formData FormLogin

	if err := c.ShouldBindWith(&formData, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	s.cfg.GLogger.Debug("Form Data:", formData)
}

// prepareInfraRoutes Method helps with route preparation.
// Routes need to contain the starting slash ex. /route.
func (s *HTTPServer) prepareLoginRoute() []route {
	routeLogin := route{
		Group:    endPointGroupAuthorization,
		Endpoint: endpointLogin,
		Method:   http.MethodPost,
		Handler:  s.handlerLogin,
	}

	return []route{routeLogin}
}
