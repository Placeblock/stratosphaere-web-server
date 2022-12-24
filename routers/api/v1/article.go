package v1

import (
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	id, interr := strconv.ParseInt(c.Param("id"), 10, 32)

	if interr != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	fields := c.QueryArray("fields")
	if len(fields) > 0 {
		fields = append(fields, "updated_at")
	}
	uid := uint16(id)

	articleSerivce := models.Article{ID: &uid}
	article, err := articleSerivce.Get(fields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appG.Response(http.StatusBadRequest, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		} else {
			appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_GET, nil)
		}
		return
	}

	timestamp := article.UpdatedAt.UTC().Format(http.TimeFormat)
	if timestamp == c.GetHeader("If-Modified-Since") {
		appG.Response(http.StatusNotModified, exception.SUCCESS, nil)
		return
	}
	c.Header("Last-Modified", timestamp)
	appG.C.Header("Cache-Control", "public, max-age=0")

	appG.Response(http.StatusOK, exception.SUCCESS, *article)
}

func GetIDChunk(c *gin.Context) {
	appG := app.Gin{C: c}

	var getArticlesParams models.GetArticlesParams

	if err := c.BindQuery(&getArticlesParams); err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	articleService := models.Article{}
  
	_, loggedIn := c.Get("user")
	if !loggedIn {
		*getArticlesParams.Unpublished = false
	}
	lastmodified, err := models.GetAllLastModified()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLES_FAIL_GET, nil)
		return
	}
	if lastmodified.Format(http.TimeFormat) == c.GetHeader("If-Modified-Since") {
		appG.Response(http.StatusNotModified, exception.SUCCESS, nil)
		return
	}
	idchunk, err := articleService.GetIDChunk(&getArticlesParams)
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLES_FAIL_GET, nil)
		return
	}
	appG.C.Header("Last-Modified", lastmodified.Format(http.TimeFormat))
	appG.C.Header("Cache-Control", "max-age=0")
	appG.Response(http.StatusOK, exception.SUCCESS, idchunk)
}

type ArticleVisibilityForm struct {
	Publish *bool `json:"publish" binding:"required"`
}

func ArticleVisibility(c *gin.Context) {
	appG := app.Gin{C: c}

	id, iderr := strconv.ParseInt(c.Param("id"), 10, 32)
	visibilityForm := ArticleVisibilityForm{}

	err := c.BindJSON(&visibilityForm)
	if iderr != nil || err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	uid := uint16(id)

	articleService := models.Article{ID: &uid}

	publishDate, err := articleService.Visibility(*visibilityForm.Publish)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appG.Response(http.StatusBadRequest, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		} else {
			appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_EDIT, nil)
		}
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, publishDate.Format(http.TimeFormat))
}

func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	author := c.GetString("user")
	articleService := models.Article{
		Author: &author,
	}
	err := articleService.Create()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CREATE, nil)
		return
	}

	article, err := articleService.Get([]string{})

	appG.Response(http.StatusOK, exception.SUCCESS, article)
}

type EditArticleForm struct {
	ID            uint16 `json:"id" binding:"required"`
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
		ID:            &form.ID,
		Title:         &form.Title,
		Description:   &form.Description,
		CoverImageUrl: &form.CoverImageUrl,
		Content:       &form.Content,
	}

	err := articleService.Edit()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appG.Response(http.StatusBadRequest, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		} else {
			appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_EDIT, nil)
		}
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, nil)
}

func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	uid := uint16(id)
	articleService := models.Article{ID: &uid}

	err := articleService.Delete()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appG.Response(http.StatusBadRequest, exception.ERROR_ARTICLE_NOT_EXIST, nil)
		} else {
			appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_DELETE, nil)
		}
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, nil)
}

func StoreImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	if header.Size >= 1000000 {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_ARTICLE_FAIL_CREATE, nil)
		return
	}
	filetype := http.DetectContentType(buff)
	if !strings.HasPrefix(filetype, "image") {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}
	c.SaveUploadedFile(header, "/home/felix/coding/web/stratosphaere/images/"+header.Filename)
	appG.Response(http.StatusOK, exception.SUCCESS, "https://stratosphaere.codelix.de/images/"+header.Filename)
}
