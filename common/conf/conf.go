package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadFromYaml(configFile string, conf interface{}) {
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	err = v.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("解析配置文件失败: %s", err))
	}
	return
}
