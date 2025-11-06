package utils

import (
	"compus-second-hand/global"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigName("config")    // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")      // 或viper.SetConfigType("YAML")
	viper.AddConfigPath("../config") // 配置文件路径
	err := viper.ReadInConfig()      // 查找并读取配置文件
	if err != nil {                  // 处理读取配置文件的错误
		panic("配置文件读取失败: " + err.Error())
	}
	if err := viper.Unmarshal(&global.Configs); err != nil {
		panic("配置文件解析失败: " + err.Error())
	}
}
