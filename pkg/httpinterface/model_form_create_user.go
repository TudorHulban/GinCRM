package httpinterface

import (
	"net/http"

	authentication "github.com/TudorHulban/GinCRM/pkg/authenticate"
	"github.com/TudorHulban/GinCRM/pkg/validator"
	"github.com/gin-gonic/gin"
)

// FormCreateUser Structure used for creating user.
type FormCreateUser struct {
	FieldUserCode string `form:"usercode" validate:"required"`
	FieldPassword string `form:"password" validate:"required"`
}

func (s *HTTPServer) handlerCreateUser(c *gin.Context) {
	var formData FormCreateUser

	// curl -X POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/createuser
	if errBind := c.Bind(&formData); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	if errValid := validator.GetValidator().Struct(formData); errValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errValid.Error()})
		return
	}

	// create user in cache
	u := authentication.UserAuth{
		Code:     formData.FieldUserCode,
		Password: formData.FieldPassword,
	}

	if errCreateCache := u.SaveToCache(); errCreateCache != nil {
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
