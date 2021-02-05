package bootstrap

import (
	"fitness/global"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var Cfg *ini.File
var err error

func InitConfig() {
	//1. 配置文件加载
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	//2.初始化
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	global.RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	global.HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	global.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	global.WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	//注意，JwtSecret必须是[]byte类型
	global.JwtSecret = []byte(sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)"))
	global.APP_NAME = sec.Key("APP_NAME").MustString("fitness")
}
