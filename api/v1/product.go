package v1

import (
	"gin_mal_tmp/pkg/util"
	"gin_mal_tmp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["image"]
	claim, _ := util.ParseToken(c.GetHeader("access_token"))
	CreateProductService := service.ProductService{}
	if err := c.ShouldBind(&CreateProductService); err == nil {
		res := CreateProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln("CreateProduct api ", err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
