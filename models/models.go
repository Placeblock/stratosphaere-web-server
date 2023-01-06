package models

import (
	"fmt"
	"log"
	"stratosphaere-server/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		//Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Auth{})
	db.AutoMigrate(&LiveData{})
}

func CloseDB() {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}
