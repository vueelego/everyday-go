package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Parser struct {
	vp *viper.Viper
}

func main() {
	fmt.Println(viper.New(), fsnotify.New())
}

func NewSettings(configs ...string) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}
