package httpinterface

import (
	"net/http"

	authentication "github.com/TudorHulban/GinCRM/pkg/logic/authenticate"
	"github.com/gin-gonic/gin"
)

// FormCreateUser Structure used for creating user.
type FormCreateUser struct {
	FieldUserCode string `form:"usercode" validate:"required"`
	FieldPassword string `form:"password" validate:"required"`
}

// Verify with:
// curl -X POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/createuser
func (s *HTTPServer) handlerCreateUser(c *gin.Context) {
	var formData FormCreateUser

	if errValid := BindAndValidate(formData, c); errValid != nil {
		return
	}

	// create user in cache
	u := authentication.UserAuth{
		Code:     formData.FieldUserCode,
		Password: formData.FieldPassword,
	}

	if errCreateCache := u.SaveToLoginCache(); errCreateCache != nil {
		c.AbortWithError(http.StatusInternalServerError, errCreateCache)
		return
	}

	c.JSON(http.StatusOK, formData)
}

// prepareLoginRoute Method helps with route preparation.
// Routes need to contain the starting slash ex. /route.
func (s *HTTPServer) prepareCreateUserRoute() []route {
	routeLogin := route{
		Group:    endPointGroupAuthorization,
		Endpoint: endPointCreateUser,
		Method:   http.MethodPost,
		Handler:  s.handlerCreateUser,
	}

	return []route{routeLogin}
}
