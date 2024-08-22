package v1

import (
	"gin_mal_tmp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	if err := c.ShouldBind(&listCarousel); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := listCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
