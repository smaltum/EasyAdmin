package bootstrap

import (
	"easy-admin/plugin/gvar"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg = new(config)

type config struct {
	cfg *viper.Viper
}

func (r *config) _init() {

	var cfgPath = "config/config.toml"

	r.cfg = viper.New()
	r.cfg.SetConfigFile(cfgPath)
	r.cfg.SetConfigType("toml")
	err := r.cfg.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	r.cfg.WatchConfig()

	r.cfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
}

func (r *config) GetVal(key string) *gvar.Var {
	return gvar.New(r.cfg.Get(key), true)
}
