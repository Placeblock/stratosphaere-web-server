package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to parse config.ini: %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
