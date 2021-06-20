package config

import (
	"encoding/json"
	"os"
	tools "restful-api-kit/utilities"
)

// Config 配置对象
type Config struct {
	Database *Database `json:"database"`
}

// GlobalConfigSetting 配置实例
var GlobalConfigSetting = &Config{}

// Setup 配置
func Setup() {
	mode := tools.GetEnv("mode")

	var filePtr  *os.File
	var err error
	if mode == "development" {
		filePtr, err = os.Open("config/dbStaging.json") //config的文件目录
	} else if mode == "production" {
		filePtr, err = os.Open("config/dbProduction.json") //config的文件目录
	} else {
		filePtr, err = os.Open("config/dbStaging.json") //config的文件目录
	}

	if err != nil {
		tools.CheckErr(err)
		return
	}
	defer func(filePtr *os.File) {
		err := filePtr.Close()
		if err != nil {
			tools.CheckErr(err)
		}
	}(filePtr)
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(GlobalConfigSetting)
	DatabaseSetting = GlobalConfigSetting.Database
}

// Database 数据库配置对象
type Database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

// DatabaseSetting 数据库配置对象 实例
var DatabaseSetting = &Database{}
