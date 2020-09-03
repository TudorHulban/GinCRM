package httpinterface

import (
	"net/http"

	authentication "github.com/TudorHulban/GinCRM/pkg/authenticate"
	"github.com/TudorHulban/GinCRM/pkg/validator"
	"github.com/gin-gonic/gin"
)

// FormLogin Structure used for validating login request.
type FormLogin struct {
	FieldUserCode string `form:"usercode" validate:"required"`
	FieldPassword string `form:"password" validate:"required"`
}

func (s *HTTPServer) handlerLogin(c *gin.Context) {
	var formData FormLogin
	// curl -X POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/login
	if errBind := c.Bind(&formData); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	if errValid := validator.GetValidator().Struct(formData); errValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errValid.Error()})
		return
	}

	s.cfg.GLogger.Debug("Form Data:", formData)

	// check if authorized. if authorized return session ID.
	// in backend insert in session cache the user structure and in user cache the credentials.
	u := authentication.UserAuth{
		Code:     formData.FieldUserCode,
		Password: formData.FieldPassword,
	}
	if errAuthenticate := u.IsAuthenticated(); errAuthenticate != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, formData)
}

// prepareLoginRoute Method helps with route preparation.
// Routes need to contain the starting slash ex. /route.
func (s *HTTPServer) prepareLoginRoute() []route {
	routeLogin := route{
		Group:    endPointGroupAuthorization,
		Endpoint: endPointLogin,
		Method:   http.MethodPost,
		Handler:  s.handlerLogin,
	}

	return []route{routeLogin}
}
