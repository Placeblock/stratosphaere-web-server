package main

import (
	"fmt"
	"log"
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/setting"
	"stratosphaere-server/pkg/util"
	"stratosphaere-server/routers"

	"github.com/gin-gonic/gin"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func init() {
	setting.Setup()
	models.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	log.Printf("Start http server on Port %s", endPoint)

	server.ListenAndServe()
}
