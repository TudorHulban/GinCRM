package httpinterface

import (
	"net/http"

	"github.com/TudorHulban/GinCRM/pkg/validator"
	"github.com/gin-gonic/gin"
)

// FormLogin Structure used for validating login request.
type FormLogin struct {
	FieldUserCode string `form:"usercode"`
	FieldPassword string `form:"password"`
}

func (s *HTTPServer) handlerLogin(c *gin.Context) {
	var formData FormLogin
	c.Bind(&formData) // curl -X POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/login

	if errData := validator.GetValidator().Struct(formData); errData != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errData.Error()})
	}

	s.cfg.GLogger.Debug("Form Data:", formData)
	c.JSON(http.StatusOK, formData)
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
