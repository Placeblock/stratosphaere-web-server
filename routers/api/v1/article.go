package v1

import (
	"fmt"
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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

	pageNum, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if err != nil {
		pageNum = 0
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 0)
	if err != nil {
		pageSize = 2147483647
	}

	articleService := models.Article{}

	articles, err := articleService.GetAll(int(pageNum), int(pageSize))
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

	data := make(map[string]interface{})
	data["articles"] = articles

	appG.Response(http.StatusOK, exception.SUCCESS, data)
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

	appG.Response(http.StatusOK, exception.SUCCESS, map[string]uint16{
		"id": articleService.ID,
	})
}

type EditArticleForm struct {
	ID            uint16 `form:"id" validate:"required"`
	Title         string `form:"title" validate:"required,max=200"`
	Description   string `form:"description" validate:"required,max=1000"`
	Content       string `form:"content" validate:"required,max=65535"`
	CoverImageUrl string `form:"cover_image_url" validate:"required,max=255"`
	Published     bool   `form:"published" validate:"required"`
}

func EditArticle(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	fmt.Println(id)

	var (
		appG = app.Gin{C: c}
		form = EditArticleForm{ID: uint16(id)}
	)

	c.Request.ParseForm()
	fmt.Printf("c.Request.Form: %v\n", c.Request.Form)
	if c.BindJSON(&form) != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	valid := validator.Validate{}
	if valid.Struct(form) != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	articleService := models.Article{
		ID:            form.ID,
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
