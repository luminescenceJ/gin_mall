package routes

import (
	api "gin_mal_tmp/api/v1"
	"gin_mal_tmp/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "success"})
		})

		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		authed := v1.Group("/") // 登录保护
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.POST("user/update", api.UserUpdate)
			authed.POST("user/avatar", api.UpdateAvatar)
			authed.POST("user/send_email", api.SendEmail)
			authed.GET("user/valid_email", api.ValidEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)

			//商品操作
			authed.POST("product/create", api.CreateProduct)
		}

	}
	return r
}
