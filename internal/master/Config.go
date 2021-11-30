package master

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Runmode      string `mapstructure:"runmode"`
	Addr         string `mapstructure:"addr"`
}

func InitConfig() error {
	viper.SetConfigFile("config/master.yaml")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("viper ReadInConfig failed, err:%v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Unmarshal to conf failed, err:%v", err))
	}
	return err
}
