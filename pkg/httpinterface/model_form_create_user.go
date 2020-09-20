package httpinterface

/*
Create user handler.
During the user creation process the cache is updated accordingly.
*/

import (
	"net/http"
	"strconv"

	"github.com/TudorHulban/GinCRM/pkg/persistence"

	authentication "github.com/TudorHulban/GinCRM/pkg/logic/authenticate"
	"github.com/gin-gonic/gin"
)

// FormCreateUser Structure used for creating user.
type FormCreateUser struct {
	FieldUserCode string `form:"usercode" validate:"required"`
	FieldPassword string `form:"password" validate:"required"`
}

// ResponseCreateUser Structure used to respond to a create user request.
type ResponseCreateUser struct {
	UserID string `json:"userID"`
}

// Verify with:
// curl -X POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/createuser
func (s *HTTPServer) handlerCreateUser(c *gin.Context) {
	var formData FormCreateUser

	if errValid := BindAndValidate(&formData, c); errValid != nil {
		return
	}

	// create user in RDBMS
	u := persistence.User{
		SecurityGroupID:   1,
		UserCode:          formData.FieldUserCode,
		PasswordLoginForm: formData.FieldPassword,
	}

	if errCreate := s.crudLogic.AddUser(&u); errCreate != nil {
		c.AbortWithError(http.StatusInternalServerError, errCreate)
		return
	}

	// create user in cache
	cache := authentication.UserAuth{
		Code:     formData.FieldUserCode,
		Password: formData.FieldPassword,
	}

	if errCreateCache := cache.SaveToLoginCache(); errCreateCache != nil {
		c.AbortWithError(http.StatusInternalServerError, errCreateCache)
		return
	}

	c.JSON(http.StatusOK, ResponseCreateUser{
		UserID: strconv.FormatUint(u.ID, 10),
	})
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
