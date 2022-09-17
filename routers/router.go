package routers

import (
	v1 "stratosphaere-server/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/v1")

	apiv1_unsecure := apiv1.Group("/")
	apiv1_unsecure.POST("/auth", v1.GetAuth)

	//UNSECURE-EMAIL
	apiv1_unsecure.POST("/email")
	apiv1_unsecure.DELETE("/email")

	//UNSECURE-LIVE
	apiv1_unsecure.GET("/live")

	//UNSECURE-BLOG
	apiv1_unsecure_blog := apiv1_unsecure.Group("/blog")

	apiv1_unsecure_blog.GET("/articles")
	apiv1_unsecure_blog.GET("/articles/:id")

	//SECURE
	apiv1_secure := apiv1.Group("/")

	//SECURE-BLOG
	apiv1_secure_blog := apiv1_secure.Group("/blog")
	apiv1_secure_blog.POST("/articles")
	apiv1_secure_blog.DELETE("/articles/:uuid")
	apiv1_secure_blog.PUT("/articles/:uuid")

	//SECURE-LIVE
	apiv1_secure.POST("/live")
}
