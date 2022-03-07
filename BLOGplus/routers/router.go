package routers

import (
	"BLOGplus/api/v1"
	"BLOGplus/middleware"
	"BLOGplus/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口

		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DeleteUser)

		//分类模块的路由接口
		Auth.POST("category/add", v1.AddCategory)
		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		Auth.POST("article/add", v1.AddArticle)

		Auth.PUT("article/:id", v1.EditArticle)
		Auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		Auth.POST("upload", v1.UpLoad)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCategory)
		router.GET("article", v1.GetArticle)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.GET("article/list/:id", v1.GetCAateArt)
		router.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
