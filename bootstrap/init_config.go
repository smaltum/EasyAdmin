package bootstrap

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Config config

func init() {
	loadConfig()
}

func loadConfig() {
	var cfgPath string = "config/config.toml"
	if _, err := toml.DecodeFile(cfgPath, &Config); err != nil {
		log.Println(err)
	}
}

type config struct {
	Server server
	Oss    oss
}

// server 服务
type server struct {
	Port               int
	Mode               string
	MaxMultipartMemory int64
}

// oss 云存储
type oss struct {
	AppKey    string
	AppSecret string
	Bucket    string
	Endpoint  string
}
