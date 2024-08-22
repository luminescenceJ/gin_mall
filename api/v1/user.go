package v1

import (
	"gin_mal_tmp/pkg/util"
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

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}

func UpdateAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size

	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&uploadAvatar); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	}
}

func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&sendEmail); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}

func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService
	if err := c.ShouldBind(&validEmail); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := validEmail.Valid(c.Request.Context(), c.Query("token"))
		c.JSON(http.StatusOK, res)
	}
}

func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&showMoney); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := showMoney.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}
