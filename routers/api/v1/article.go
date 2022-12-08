package v1

import (
	"fmt"
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	articleSerivce := models.Article{ID: uint16(id)}
	exists, err := articleSerivce.Exists()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CHECK_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusBadRequest, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		return
	}

	article, err := articleSerivce.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_GET, nil)
		return
	}
	appG.Response(http.StatusOK, exception.SUCCESS, article)
}

type GetArticlesParams struct {
	Offset    *uint16 `form:"offset" json:"offset"  binding:"required"`
	Limit     *uint16 `form:"limit" json:"limit"  binding:"required"`
	Published *bool   `form:"published" json:"published"`
}

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}

	var getArticlesParams GetArticlesParams

	if err := c.BindQuery(&getArticlesParams); err != nil {
		fmt.Println(err)
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	articleService := models.Article{}

	_, loggedIn := c.Get("user")

	articles, err := articleService.GetAll(int(*getArticlesParams.Offset), int(*getArticlesParams.Limit), !loggedIn || (getArticlesParams.Published != nil && *getArticlesParams.Published))

	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_COUNT, nil)
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, articles)
}

type ArticleVisibilityForm struct {
	Publish *bool `json:"publish" binding:"required"`
}

func ArticleVisibility(c *gin.Context) {
	appG := app.Gin{C: c}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	visibilityForm := ArticleVisibilityForm{}

	err := c.BindJSON(&visibilityForm)
	if err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	articleService := models.Article{
		ID: uint16(id),
	}

	exists, err := articleService.Exists()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CHECK_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusBadRequest, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		return
	}

	publishDate, err := articleService.Visibility(*visibilityForm.Publish)
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_EDIT, nil)
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, publishDate)
}

func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	author := c.GetString("user")
	articleService := models.Article{
		ArticleMetadata: models.ArticleMetadata{Author: author},
	}
	err := articleService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CREATE, nil)
		return
	}

	article, err := articleService.Get()

	appG.Response(http.StatusOK, exception.SUCCESS, article)
}

type EditArticleForm struct {
	ID            int    `json:"id" binding:"required"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
}

func EditArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	form := EditArticleForm{}

	if c.BindJSON(&form) != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	articleService := models.Article{
		ID: uint16(form.ID),
		ArticleMetadata: models.ArticleMetadata{
			Title:         form.Title,
			Description:   form.Description,
			CoverImageUrl: form.CoverImageUrl},
		Content: form.Content,
	}

	exists, err := articleService.Exists()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CHECK_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_EDIT, nil)
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, nil)
}

func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	articleService := models.Article{ID: uint16(id)}
	exists, err := articleService.Exists()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CHECK_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_DELETE, nil)
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, nil)
}
