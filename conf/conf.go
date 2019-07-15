package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	// Cfg .
	Cfg *ini.File
)

// 加载配置文件
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("Loade config err")
	}
	log.Println(Cfg)
}

