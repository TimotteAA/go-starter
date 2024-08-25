package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type ConfigFunc = func() interface{}

var (
	// viper实例
	viperInstance *viper.Viper
	configFuncs map[string]ConfigFunc
)

func init() {
	viperInstance = viper.New();
	// 从env文件中读取配置内容
	viperInstance.SetConfigType("env")
	// 配置文件和main.go同级
	viper.AddConfigPath(".")
	// 指定.env中key的前缀
	viper.SetEnvPrefix("APP_ENV")
	// 读取环境变量，包含命令行参数
	viper.AutomaticEnv()

	configFuncs = make(map[string]func() interface{})
}


// 读取配置文件
func InitConfig(mode string) {
	var (
		configFilePath = ".env" + strings.Trim(mode, ".")
	)

	viper.New()
	
	// 读取.env文件
	viper.SetConfigType(".env")
	// 相对于main.go而言 .env.xxx
	viper.AddConfigPath(".")
	viper.AddConfigPath(configFilePath) // 也可以添加多个路径

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 把配置函数中的内容放到viper中
	for key, configFunc := range configFuncs {
		viper.Set(key, configFunc())
	}
}

// Add 函数将一个 ConfigFunc 添加到 configFuncs 中
func Add(name string, fn ConfigFunc) {
	configFuncs[name] = fn
}

func init() {

}