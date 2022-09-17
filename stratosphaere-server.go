package main

import (
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
	readTimeout := setting
}
