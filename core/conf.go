package core

import (
	"fmt"
	"gin-box/conf"
	"gin-box/global"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// InitCore 读取 yaml 文件的配置
func InitCore() {
	const ConfigFile = "settings.yaml"
	c := &conf.Config{}
	yamlConfFile, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败：%s", err))
	}
	err = yaml.Unmarshal(yamlConfFile, c)
	if err != nil {
		log.Fatalf("yaml文件转换错误：%v", err)
	}
	log.Println("配置文件初始化成功！")
	global.Config = c
}
