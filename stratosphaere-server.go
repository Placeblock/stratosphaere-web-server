package main

import (
	"fmt"
	"log"
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/setting"
	"stratosphaere-server/pkg/util"
	"stratosphaere-server/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	util.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("Start http server on Port %s", endPoint)

	server.ListenAndServe()
}
