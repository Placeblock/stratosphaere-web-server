package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLiveData(c *gin.Context) {
	appG := app.Gin{C: c}

	var getLiveDataParams models.GetLiveDataParams

	if err := c.BindQuery(&getLiveDataParams); err != nil {
		fmt.Println(err)
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	liveData, err := models.GetLiveData(*getLiveDataParams.Since)

	if err != nil {
		appG.Response(http.StatusBadRequest, exception.ERROR_LIVE_DATA_GET, nil)
		return
	}

	fmt.Println(liveData)
	appG.Response(http.StatusOK, exception.SUCCESS, liveData)
}

func SetLiveData(c *gin.Context) {
	jsonData, _ := c.GetRawData()
	var liveData models.SetLiveData
	json.Unmarshal(jsonData, &liveData)
	nowTime := time.Now()
	liveData.UplinkMessage.DecodedPayload.CreatedAt = &nowTime
	liveData.UplinkMessage.DecodedPayload.Create()
	c.Status(http.StatusOK)
}
