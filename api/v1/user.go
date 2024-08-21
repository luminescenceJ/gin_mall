package v1

import (
	"gin_mal_tmp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
