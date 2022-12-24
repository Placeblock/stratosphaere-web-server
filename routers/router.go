package routers

import (
	"net/http"
	"stratosphaere-server/middleware"
	v1 "stratosphaere-server/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	//r.Use(gin.Logger())
	r.Use(middleware.CORS())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Working!")
	})

	apiv1 := r.Group("/v1")
	apiv1.Use(middleware.JWT())

	apiv1_unsecure := apiv1.Group("")
	apiv1_unsecure.POST("/auth", v1.GetAuth)

	//UNSECURE-EMAIL
	/*apiv1_unsecure.POST("/email")
	apiv1_unsecure.DELETE("/email")*/

	//UNSECURE-LIVE
	/*apiv1_unsecure.GET("/live")*/

	//UNSECURE-BLOG
	apiv1_unsecure_blog := apiv1_unsecure.Group("/blog")

	apiv1_unsecure_blog.GET("/articles", v1.GetIDChunk)
	apiv1_unsecure_blog.GET("/articles/:id", v1.GetArticle)

	//SECURE
	apiv1_secure := apiv1.Group("")
	apiv1_secure.Use(middleware.Restrict())

	//SECURE-BLOG
	apiv1_secure_blog := apiv1_secure.Group("/blog")
	apiv1_secure_blog.POST("/articles", v1.AddArticle)
	apiv1_secure_blog.DELETE("/articles/:id", v1.DeleteArticle)
	apiv1_secure_blog.PUT("/articles/:id", v1.EditArticle)
	apiv1_secure_blog.PUT("/articles/:id/publish", v1.ArticleVisibility)
	apiv1_secure_blog.POST("/image", v1.StoreImage)

	//SECURE-LIVE
	/*apiv1_secure.POST("/live")*/

	return r
}
