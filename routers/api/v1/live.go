package v1

import (
	"fmt"
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"

	"github.com/gin-gonic/gin"
)

func GetLiveData(c *gin.Context) {
	appG := app.Gin{C: c}

	var getLiveDataParams models.GetLiveDataParams

	if err := c.BindQuery(&getLiveDataParams); err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	fmt.Println(getLiveDataParams)
}
