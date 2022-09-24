package v1

import (
	"fmt"
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"stratosphaere-server/pkg/util"
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
		appG.Response(http.StatusOK, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		return
	}

	article, err := articleSerivce.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_GET, nil)
		return
	}
	appG.Response(http.StatusOK, exception.SUCCESS, article)
}

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 32)
	if err != nil {
		offset = 0
	}
	amount, err := strconv.ParseInt(c.Query("amount"), 10, 0)
	if err != nil {
		amount = 2147483647
	}

	articleService := models.Article{}

	articles, err := articleService.GetAll(int(offset), int(amount))
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_COUNT, nil)
		return
	}

	_, exists := c.Get("user")
	if !exists {
		published := []*models.Article{}
		for i := range articles {
			if articles[i].Published {
				published = append(published, articles[i])
			}
		}
		articles = published
	}

	appG.Response(http.StatusOK, exception.SUCCESS, articles)
}

func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	author := c.GetString("user")
	articleService := models.Article{
		Author: author,
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
	ID            int    `json:"id" binding:"required" validate:"required"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	Published     bool   `json:"published"`
}

func EditArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	form := EditArticleForm{}

	if c.BindJSON(&form) != nil {
		fmt.Println("BIND JSON ERROR")
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	err := util.Validate.Struct(&form)
	if err != nil {
		fmt.Println("VALIDATE STRUCT ERROR")
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	articleService := models.Article{
		ID:            uint16(form.ID),
		Title:         form.Title,
		Description:   form.Description,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		Published:     form.Published,
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
