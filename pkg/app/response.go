package app

import (
	"stratosphaere-server/pkg/exception"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code     int         `json:"code"`
	HttpCode int         `json:"http_code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:     errCode,
		HttpCode: httpCode,
		Msg:      exception.GetMsg(errCode),
		Data:     data,
	})
}
