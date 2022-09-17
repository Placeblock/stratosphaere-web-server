package v1

import (
	"fmt"
	"net/http"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	article_serivce "stratosphaere-server/service/article_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	articleSerivce := article_serivce.Article{ID: uint16(id)}
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
	fmt.Println("GET ARTICLES")
	appG := app.Gin{C: c}

	pageNum, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 0)
	if err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	articleService := article_serivce.Article{
		PageNum:  int(pageNum),
		PageSize: int(pageSize),
	}

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_COUNT, nil)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_COUNT, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, exception.SUCCESS, data)
}

type AddArticleForm struct {
	Author string
}

func AddArticle(c *gin.Context) {
	fmt.Println("ADD ARTICLE")
	appG := app.Gin{C: c}

	author := c.GetString("user")

	articleService := article_serivce.Article{
		Author: author,
	}
	id, err := articleService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CREATE, nil)
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, map[string]uint16{
		"id": id,
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

	var (
		appG = app.Gin{C: c}
		form = EditArticleForm{ID: uint16(id)}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != exception.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleService := article_serivce.Article{
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

	articleService := article_serivce.Article{ID: uint16(id)}
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
